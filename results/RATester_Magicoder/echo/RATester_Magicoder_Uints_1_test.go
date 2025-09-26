package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUints_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			switch sourceParam {
			case "ids":
				return []string{"1", "2", "3"}
			default:
				return nil
			}
		},
	}

	var dest []uint
	err := b.Uints("ids", &dest).BindError()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []uint{1, 2, 3}
	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("Expected %v, but got %v", expected, dest)
	}
}
