package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlicePad_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice := []interface{}{1, 2, 3}
	size := 5
	val := 0
	want := []interface{}{1, 2, 3, 0, 0}
	got := SlicePad(slice, size, val)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SlicePad() = %v, want %v", got, want)
	}
}
