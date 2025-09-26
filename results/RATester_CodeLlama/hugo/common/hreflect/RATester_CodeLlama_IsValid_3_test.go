package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsValid_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var v reflect.Value
	if IsValid(v) {
		t.Errorf("IsValid(%v) = true, want false", v)
	}

	v = reflect.ValueOf(1)
	if !IsValid(v) {
		t.Errorf("IsValid(%v) = false, want true", v)
	}
}
