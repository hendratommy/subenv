package envsource

import "testing"

func TestComposeSources(t *testing.T) {
	composed := ComposeSources(&File{}, &OS{})
	for i, source := range composed.sources {
		switch source.(type) {
		case *File:
			if i != 0 {
				t.Errorf("File should be in index 0")
			}
		case *OS:
			if i != 1 {
				t.Errorf("OS should be in index 1")
			}
		default:
			t.Errorf("unrecongnized envsource type")
		}
	}
}

func TestComposed_LookupEnv(t *testing.T) {
	var fileEnv = &File{}
	if err := fileEnv.Load("test.env"); err != nil {
		t.Fatalf("failed to load test.env: %v", err)
	}

	composedEnv := ComposeSources(fileEnv, &OS{})
	if v, ok := composedEnv.LookupEnv("PATH"); !ok {
		t.Errorf("PATH environment is not set")
	} else {
		if v == "" {
			t.Errorf("PATH environment is empty")
		}
	}
	if v, ok := composedEnv.LookupEnv("NAME"); !ok {
		t.Errorf("NAME environment is not set: %s", fileEnv.env)
	} else {
		if v != "putenv" {
			t.Errorf("NAME environment should equals putenv: %s", v)
		}
	}
}
