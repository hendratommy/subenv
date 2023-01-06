package envsource

import "testing"

func TestOS_LookupEnv(t *testing.T) {
	var osEnv = &OS{}

	if v, ok := osEnv.LookupEnv("PATH"); !ok {
		t.Errorf("PATH environment is not set")
	} else {
		if v == "" {
			t.Errorf("PATH environment is empty")
		}
	}

	if _, ok := osEnv.LookupEnv("NOOP"); ok {
		t.Errorf("NOOP environment should not be set")
	}
}
