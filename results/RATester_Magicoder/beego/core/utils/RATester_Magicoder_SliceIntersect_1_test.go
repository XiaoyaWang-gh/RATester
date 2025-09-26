package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceIntersect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice1 := []interface{}{1, 2, 3, 4, 5}
	slice2 := []interface{}{4, 5, 6, 7, 8}
	expected := []interface{}{4, 5}
	result := SliceIntersect(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
