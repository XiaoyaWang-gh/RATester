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

	indexes := []string{"index1", "index2"}
	hint := IgnoreIndex(indexes...)

	if hint.GetKey() != KeyIgnoreIndex {
		t.Errorf("Expected key to be %v, but got %v", KeyIgnoreIndex, hint.GetKey())
	}

	expectedValue := indexes
	actualValue := hint.GetValue()

	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Errorf("Expected value to be %v, but got %v", expectedValue, actualValue)
	}
}
