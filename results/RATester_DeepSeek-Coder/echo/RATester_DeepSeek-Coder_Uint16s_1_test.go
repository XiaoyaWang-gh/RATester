package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUint16s_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			if sourceParam == "ids" {
				return []string{"1", "2"}
			}
			return nil
		},
	}

	var dest []uint16
	err := b.Uint16s("ids", &dest).BindError()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []uint16{1, 2}
	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("Expected %v, got %v", expected, dest)
	}
}
