package config

import (
	"fmt"
	"testing"
)

func TestType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &BoolFuncFlag{}
	if f.Type() != "bool" {
		t.Errorf("f.Type() = %v, want %v", f.Type(), "bool")
	}
}
