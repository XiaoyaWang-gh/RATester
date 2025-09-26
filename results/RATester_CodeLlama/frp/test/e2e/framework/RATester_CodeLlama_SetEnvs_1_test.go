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

	f := NewDefaultFramework()
	envs := []string{"key1=value1", "key2=value2"}
	f.SetEnvs(envs)
	if !reflect.DeepEqual(envs, f.osEnvs) {
		t.Errorf("expect %v, but got %v", envs, f.osEnvs)
	}
}
