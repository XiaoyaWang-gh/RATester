package math

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttoFloatsE_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Converting a slice of numbers to floats
	slice := []any{1, 2, 3, 4, 5}
	expectedFloats := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	floats, _, _ := ns.toFloatsE(slice)
	if !reflect.DeepEqual(floats, expectedFloats) {
		t.Errorf("Expected %v, but got %v", expectedFloats, floats)
	}

	// Test case 2: Converting a single number to a float
	single := 10
	expectedSingle := []float64{10.0}
	singleFloat, _, _ := ns.toFloatsE(single)
	if !reflect.DeepEqual(singleFloat, expectedSingle) {
		t.Errorf("Expected %v, but got %v", expectedSingle, singleFloat)
	}

	// Test case 3: Converting a non-number to a float
	nonNumber := "hello"
	_, _, err := ns.toFloatsE(nonNumber)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
