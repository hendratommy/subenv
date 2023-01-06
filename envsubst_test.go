package subenv

import (
	"strings"
	"testing"

	"github.com/hendratommy/subenv/codec"
	"github.com/hendratommy/subenv/envsource"
)

func TestEnvSubst_findVariables(t *testing.T) {
	var testVars = `$NAME
$name
noop
${PREFIX}name
name${SUFFIX}
'$QUOTES'
"$DOUBLE_QUOTES"
$HELLO1$WORLD1
$HELLO2${WORLD2}
${HELLO3}$WORLD3
${HELLO4}${WORLD4}
$HELLO5 ALT $WORLD5
`
	env := &EnvSubst{content: testVars}
	res := env.findVariables()
	if len(res) != 16 {
		t.Errorf("expect len %d got %d\n", 6, len(res))
	}
	for i, val := range res {
		switch i {
		case 0:
			if val != "$NAME" {
				t.Errorf("%s should equals $NAME", val)
			}
		case 1:
			if val != "$name" {
				t.Errorf("%s should equals $name", val)
			}
		case 2:
			if val != "${PREFIX}" {
				t.Errorf("%s should equals ${PREFIX}", val)
			}
		case 3:
			if val != "${SUFFIX}" {
				t.Errorf("%s should equals ${SUFFIX}", val)
			}
		case 4:
			if val != "$QUOTES" {
				t.Errorf("%s should equals $QUOTES", val)
			}
		case 5:
			if val != "$DOUBLE_QUOTES" {
				t.Errorf("%s should equals $DOUBLE_QUOTES", val)
			}
		case 6:
			if val != "$HELLO1" {
				t.Errorf("%s should equals $HELLO1", val)
			}
		case 7:
			if val != "$WORLD1" {
				t.Errorf("%s should equals $WORLD1", val)
			}
		case 8:
			if val != "$HELLO2" {
				t.Errorf("%s should equals $HELLO2", val)
			}
		case 9:
			if val != "${WORLD2}" {
				t.Errorf("%s should equals ${WORLD2}", val)
			}
		case 10:
			if val != "${HELLO3}" {
				t.Errorf("%s should equals ${HELLO3}", val)
			}
		case 11:
			if val != "$WORLD3" {
				t.Errorf("%s should equals $WORLD3", val)
			}
		case 12:
			if val != "${HELLO4}" {
				t.Errorf("%s should equals ${HELLO4}", val)
			}
		case 13:
			if val != "${WORLD4}" {
				t.Errorf("%s should equals ${WORLD4}", val)
			}
		case 14:
			if val != "$HELLO5" {
				t.Errorf("%s should equals $HELLO5", val)
			}
		case 15:
			if val != "$WORLD5" {
				t.Errorf("%s should equals $WORLD5", val)
			}
		}
	}
}

func TestEnvSubst_Substitute(t *testing.T) {
	var testVars = `$NAME
$name
noop
${PREFIX}name
name${SUFFIX}
'$QUOTES'
"$DOUBLE_QUOTES"
$HELLO1$WORLD1
$HELLO2${WORLD2}
${HELLO3}$WORLD3
${HELLO4}${WORLD4}
$HELLO5 ALT $WORLD5
$NOT_FOUND
`
	fileEnv := &envsource.File{}
	fileEnv.Load("test.env")
	env := NewEnvSubstBuilder(testVars).SetEnvSource(fileEnv).Build()
	res, err := env.Substitute()
	if err != nil {
		t.Fatalf("substitute should not return error")
	}
	lines := strings.Split(strings.ReplaceAll(res, "\r\n", "\n"), "\n")
	for i, line := range lines {
		switch i {
		case 0:
			if line != "SUBENV" {
				t.Errorf("%s should equals SUBENV", line)
			}
		case 1:
			if line != "subenv" {
				t.Errorf("%s should equals subenv", line)
			}
		case 2:
			if line != "noop" {
				t.Errorf("%s should equals noop", line)
			}
		case 3:
			if line != "prefix-name" {
				t.Errorf("%s should equals prefix-name", line)
			}
		case 4:
			if line != "name-suffix" {
				t.Errorf("%s should equals name-suffix", line)
			}
		case 5:
			if line != "'lorem'" {
				t.Errorf("%s should equals 'lorem'", line)
			}
		case 6:
			if line != `"ipsum"` {
				t.Errorf(`%s should equals "ipsum"`, line)
			}
		case 7:
			if line != "hello1world1" {
				t.Errorf("%s should equals hello1world1", line)
			}
		case 8:
			if line != "hello2world2" {
				t.Errorf("%s should equals hello2world2", line)
			}
		case 9:
			if line != "hello3world3" {
				t.Errorf("%s should equals hello3world3", line)
			}
		case 10:
			if line != "hello4world4" {
				t.Errorf("%s should equals hello4world4", line)
			}
		case 11:
			if line != "hello5 ALT world5" {
				t.Errorf("%s should equals hello5 ALT world5", line)
			}
		case 12:
			if line != "$NOT_FOUND" {
				t.Errorf("%s should equals $NOT_FOUND", line)
			}
		}
	}
}

func TestEnvSubst_Substitute_Decode(t *testing.T) {
	var testVars = "$B64"
	fileEnv := &envsource.File{}
	fileEnv.Load("test.env")
	env := NewEnvSubstBuilder(testVars).
		SetEnvSource(fileEnv).
		SetDecoder(&codec.Base64Codec{}).
		Build()

	res, err := env.Substitute()
	if err != nil {
		t.Fatalf("substitute should not return error")
	}
	if res != "go-envsubst" {
		t.Errorf("%s should equals go-envsubst", res)
	}

	testVars = "$NAME"
	env = NewEnvSubstBuilder(testVars).
		SetEnvSource(fileEnv).
		SetDecoder(&codec.Base64Codec{}).
		Build()

	res, err = env.Substitute()
	if err == nil {
		t.Fatalf("substitute should return error")
	}
	if res != "" {
		t.Errorf("should empty, got: %s", res)
	}
}

func TestEnvSubst_Substitute_Encode(t *testing.T) {
	var testVars = "$NAME"
	fileEnv := &envsource.File{}
	fileEnv.Load("test.env")
	env := NewEnvSubstBuilder(testVars).
		SetEnvSource(fileEnv).
		SetEncoder(&codec.Base64Codec{}).
		Build()

	res, err := env.Substitute()
	if err != nil {
		t.Fatalf("substitute should not return error")
	}
	if res != "U1VCRU5W" {
		t.Errorf("%s should equals U1VCRU5W", res)
	}
}
