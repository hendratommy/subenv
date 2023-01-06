package subenv

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hendratommy/subenv/envsource"
)

var regexPattern = `(\${[^\}]\w+})|(\$\w+)`

type Source interface {
	LookupEnv(k string) (string, bool)
}

type Encoder interface {
	Encode(s string) string
}

type Decoder interface {
	Decode(s string) (string, error)
}

type EnvSubst struct {
	content string
	env     envsource.Source
	enc     Encoder
	dec     Decoder
}

func NewEnvSubst(content string) *EnvSubst {
	return &EnvSubst{content: content, env: &envsource.OS{}}
}

func (e *EnvSubst) findVariables() []string {
	return regexp.MustCompile(regexPattern).FindAllString(e.content, -1)
}

func (e *EnvSubst) Substitute() (string, error) {
	vars := e.findVariables()

	for _, k := range vars {
		key := k
		if key[:2] == "${" {
			key = key[2 : len(key)-1]
		} else {
			key = key[1:]
		}
		if v, ok := e.env.LookupEnv(key); ok {
			if e.dec != nil {
				v1, err := e.dec.Decode(v)
				if err != nil {
					return "", fmt.Errorf("error: failed to decode %s: %v", v, err)
				}
				v = v1
			}
			if e.enc != nil {
				v = e.enc.Encode(v)
			}
			e.substitute(k, v)
		}
	}

	return e.content, nil
}

func (e *EnvSubst) substitute(old string, new string) {
	e.content = strings.ReplaceAll(e.content, old, new)
}

type EnvSubstBuilder struct {
	envSubst *EnvSubst
}

func NewEnvSubstBuilder(content string) *EnvSubstBuilder {
	return &EnvSubstBuilder{envSubst: NewEnvSubst(content)}
}

func (b *EnvSubstBuilder) SetEncoder(enc Encoder) *EnvSubstBuilder {
	b.envSubst.enc = enc
	return b
}

func (b *EnvSubstBuilder) SetDecoder(dec Decoder) *EnvSubstBuilder {
	b.envSubst.dec = dec
	return b
}

func (b *EnvSubstBuilder) SetEnvSource(env envsource.Source) *EnvSubstBuilder {
	b.envSubst.env = env
	return b
}

func (b *EnvSubstBuilder) Build() *EnvSubst {
	return b.envSubst
}
