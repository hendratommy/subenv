package envsource

import "testing"

func TestFile_LookupEnv(t *testing.T) {
	var fileEnv = &File{}

	if err := fileEnv.Load("test.env"); err != nil {
		t.Errorf("error: failed to load env files:%v", err)
	} else {
		if v, ok := fileEnv.LookupEnv("NAME"); !ok {
			t.Errorf("NAME environment is not set: %s", fileEnv.env)
		} else {
			if v != "putenv" {
				t.Errorf("NAME environment should equals 'putenv': %s", v)
			}
		}
		if _, ok := fileEnv.LookupEnv("PATH"); ok {
			t.Errorf("PATH environment should not be set")
		}
	}
}
