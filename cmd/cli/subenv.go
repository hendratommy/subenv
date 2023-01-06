package main

import (
	"flag"
	"fmt"
	"github.com/hendratommy/subenv"
	"log"
	"os"
	"path/filepath"

	"github.com/hendratommy/subenv/codec"
	"github.com/hendratommy/subenv/envsource"
)

type multiValues []string

func (e *multiValues) String() string {
	return "Environment variable files"
}

func (e *multiValues) Set(s string) error {
	*e = append(*e, s)
	return nil
}

func (e *multiValues) Size() int {
	return len(*e)
}

func validateFile(f string) error {
	if stat, err := os.Stat(f); err == nil {
		if stat.IsDir() {
			return fmt.Errorf("error: %s is a directory", f)
		}
		return nil
	}
	return fmt.Errorf("error: %s not found", f)
}

func mapEncoderName(e string) (subenv.Encoder, error) {
	switch e {
	case "base64":
		return &codec.Base64Codec{}, nil
	default:
		return nil, fmt.Errorf("error: unrecognized encoder name %s", e)
	}
}

func mapDecoderName(e string) (subenv.Decoder, error) {
	switch e {
	case "base64":
		return &codec.Base64Codec{}, nil
	default:
		return nil, fmt.Errorf("error: unrecognized decoder name %s", e)
	}
}

func composeEnvSources(envFiles multiValues, noos bool) envsource.Source {
	if envFiles.Size() > 0 {
		fileEnv := &envsource.File{}
		if err := fileEnv.Load(envFiles[0], envFiles[1:]...); err != nil {
			log.Fatalln(err)
		}
		if noos {
			return fileEnv
		}
		// prioritize files over OS
		return envsource.ComposeSources(fileEnv, &envsource.OS{})
	}
	return &envsource.OS{}
}

var (
	version string
	goos    string
	goarch  string
)

func printInfo() {
	fmt.Printf("subenv version %s %s/%s\n", version, goos, goarch)
}

func main() {
	var (
		envFiles     multiValues
		noOS         bool
		encName      string
		decName      string
		printVersion bool

		enc subenv.Encoder
		dec subenv.Decoder

		subst *subenv.EnvSubst
	)
	flag.Var(&envFiles, "e", "File to use as env source")
	flag.BoolVar(&noOS, "noos", false, `This flag can only be use when using -e. If set to true, will 
not use OS environment variables`)
	flag.StringVar(&encName, "c", "", "Encoder name to use, available encoder: [ base64 ]")
	flag.StringVar(&decName, "d", "", "Decoder name to use, available decoder: [ base64 ]")
	flag.BoolVar(&printVersion, "v", false, "Print version")
	flag.BoolVar(&printVersion, "version", false, "Print version")
	flag.Parse()
	args := flag.Args()

	if printVersion {
		printInfo()
		os.Exit(0)
	}

	if len(args) != 1 {
		log.Fatalf("Got %d as argument, requires exactly 1 argument as input file. Use \"subenv --help\" for help\n", len(args))
	}

	file, err := filepath.Abs(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	if encName != "" {
		if enc1, err := mapEncoderName(encName); err != nil {
			log.Fatalln(err)
		} else {
			enc = enc1
		}

	}
	if decName != "" {
		if dec1, err := mapDecoderName(decName); err != nil {
			log.Fatalln(err)
		} else {
			dec = dec1
		}
	}

	env := composeEnvSources(envFiles, noOS)
	if err := validateFile(file); err != nil {
		log.Fatalln(err)
	}
	if content, err := os.ReadFile(file); err != nil {
		log.Fatalln(err)
	} else {
		substBuilder := subenv.NewEnvSubstBuilder(string(content))
		substBuilder.SetEnvSource(env)
		if enc != nil {
			substBuilder.SetEncoder(enc)
		}
		if dec != nil {
			substBuilder.SetDecoder(dec)
		}
		subst = substBuilder.Build()
	}

	out, err := subst.Substitute()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(out)
}
