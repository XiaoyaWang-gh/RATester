package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniqueNonEmptyStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := []string{"", "a", "a", "b", "b", "c", "c", "d", "d", "e", "e", "f", "f", "g", "g", "h", "h", "i", "i", "j", "j", "k", "k", "l", "l", "m", "m", "n", "n", "o", "o", "p", "p", "q", "q", "r", "r", "s", "s", "t", "t", "u", "u", "v", "v", "w", "w", "x", "x", "y", "y", "z", "z"}
	expected := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	actual := uniqueNonEmptyStrings(s)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("uniqueNonEmptyStrings(%v) = %v; expected %v", s, actual, expected)
	}
}
