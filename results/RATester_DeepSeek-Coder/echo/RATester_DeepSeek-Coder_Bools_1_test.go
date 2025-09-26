package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBools_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "boolParam" {
				return "true"
			}
			return ""
		},
	}

	var dest []bool
	err := b.Bools("boolParam", &dest).BindError()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := []bool{true}
	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("Expected %v, got %v", expected, dest)
	}
}
