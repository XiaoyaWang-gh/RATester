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
		Vars: map[string]string{},
	}
	key := "key"
	val := "val"
	envVar.set(key, val)
	if envVar.Vars[key] != val {
		t.Errorf("envVar.Vars[key] = %v, want %v", envVar.Vars[key], val)
	}
}
