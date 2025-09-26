package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewResourceTransformationKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "testName"
	elements := []any{1, "test", true}
	key := NewResourceTransformationKey(name, elements...)

	if key.Name != name {
		t.Errorf("Expected Name to be %s, but got %s", name, key.Name)
	}

	if !reflect.DeepEqual(key.elements, elements) {
		t.Errorf("Expected elements to be %v, but got %v", elements, key.elements)
	}
}
