package envsource

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	env map[string]string
}

func (e *File) Load(file string, files ...string) error {
	if e.env == nil {
		e.env = make(map[string]string)
	}

	files = prepend(files, file)

	for _, f := range files {
		f1, err := filepath.Abs(f)
		if err != nil {
			return err
		}
		if err = e.load(f1); err != nil {
			return err
		}
	}
	return nil
}

func (e *File) load(f string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Err() == nil {
		line := scanner.Text()
		if index := strings.Index(line, "="); index >= 0 {
			if key := strings.TrimSpace(line[:index]); len(key) > 0 {
				value := ""
				if len(line) > index {
					value = strings.TrimSpace(line[index+1:])
				}
				e.env[key] = value
			}
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}

func (e *File) LookupEnv(k string) (string, bool) {
	if v, ok := e.env[k]; ok {
		return v, true
	}
	return "", false
}
