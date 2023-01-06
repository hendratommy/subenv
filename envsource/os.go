package envsource

import "os"

type OS struct{}

// LookupEnv retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise, the returned value will be empty and the boolean will
// be false.
func (e *OS) LookupEnv(k string) (string, bool) {
	return os.LookupEnv(k)
}
