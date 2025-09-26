package echo

import (
	"fmt"
	"testing"
)

func TestUint8s_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "123"
			}
			return ""
		},
	}

	var dest []uint8
	err := b.Uint8s("test", &dest).BindError()
	if err != nil {
		t.Errorf("Uint8s() error = %v", err)
		return
	}

	if len(dest) != 1 || dest[0] != 123 {
		t.Errorf("Uint8s() = %v, want %v", dest, []uint8{123})
	}
}
