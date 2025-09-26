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
			if sourceParam == "ids" {
				return []string{"1", "2", "3"}
			}
			return nil
		},
	}

	var dest []uint
	err := b.Uints("ids", &dest).BindError()
	if err != nil {
		t.Errorf("Uints() error = %v", err)
		return
	}

	expected := []uint{1, 2, 3}
	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("Uints() = %v, want %v", dest, expected)
	}
}
