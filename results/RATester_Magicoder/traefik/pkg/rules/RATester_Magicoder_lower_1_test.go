package rules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLower_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testSlice := []string{"HELLO", "WORLD", "GO"}
	expected := []string{"hello", "world", "go"}
	result := lower(testSlice)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
