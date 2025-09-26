package config

import (
	"fmt"
	"testing"
)

func TestString_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &BoolFuncFlag{v: true}
	if f.String() != "true" {
		t.Errorf("f.String() = %s, want %s", f.String(), "true")
	}
}
