package hints

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIgnoreIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hint := IgnoreIndex("index1", "index2")

	if hint.GetKey() != KeyIgnoreIndex {
		t.Errorf("Expected key to be %v, got %v", KeyIgnoreIndex, hint.GetKey())
	}

	expectedValue := []string{"index1", "index2"}
	if !reflect.DeepEqual(hint.GetValue(), expectedValue) {
		t.Errorf("Expected value to be %v, got %v", expectedValue, hint.GetValue())
	}
}
