package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTemplateNameAndVariants_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	expectedName := "test"
	expectedVariants := []string{"test"}
	actualName, actualVariants := templateNameAndVariants(name)
	if actualName != expectedName {
		t.Errorf("Expected name %s, got %s", expectedName, actualName)
	}
	if !reflect.DeepEqual(actualVariants, expectedVariants) {
		t.Errorf("Expected variants %v, got %v", expectedVariants, actualVariants)
	}
}
