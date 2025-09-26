package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceDiff_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice1 := []interface{}{1, 2, 3, 4}
	slice2 := []interface{}{3, 4, 5, 6}
	expected := []interface{}{1, 2}
	result := SliceDiff(slice1, slice2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
