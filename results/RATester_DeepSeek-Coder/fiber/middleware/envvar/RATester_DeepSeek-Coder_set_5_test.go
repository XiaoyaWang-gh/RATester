package envvar

import (
	"fmt"
	"testing"
)

func TestSet_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	envVar := &EnvVar{
		Vars: make(map[string]string),
	}

	key := "testKey"
	val := "testVal"

	envVar.set(key, val)

	if envVar.Vars[key] != val {
		t.Errorf("Expected %s to be %s, got %s", key, val, envVar.Vars[key])
	}
}
