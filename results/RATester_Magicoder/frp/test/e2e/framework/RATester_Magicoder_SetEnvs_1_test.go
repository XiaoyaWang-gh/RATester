package framework

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetEnvs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	envs := []string{"key1=value1", "key2=value2"}
	f := &Framework{}
	f.SetEnvs(envs)

	if !reflect.DeepEqual(f.osEnvs, envs) {
		t.Errorf("Expected envs to be %v, but got %v", envs, f.osEnvs)
	}
}
