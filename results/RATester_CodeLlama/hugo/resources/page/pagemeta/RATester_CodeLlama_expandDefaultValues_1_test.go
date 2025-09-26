package pagemeta

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExpandDefaultValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	values := []string{"a", "b", "c", ":default"}
	defaults := []string{"d", "e", "f"}
	expected := []string{"a", "b", "c", "d", "e", "f"}
	actual := expandDefaultValues(values, defaults)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expandDefaultValues(%v, %v) = %v; expected %v", values, defaults, actual, expected)
	}
}
