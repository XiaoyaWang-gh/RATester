package hydratation

import (
	"fmt"
	"testing"
)

func TestHydrate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string
		Field2 int
	}

	testStruct := TestStruct{}
	err := Hydrate(&testStruct)
	if err != nil {
		t.Errorf("Hydrate() returned an error: %v", err)
	}

	if testStruct.Field1 != "value1" {
		t.Errorf("Hydrate() did not set Field1 correctly, expected 'value1', got %v", testStruct.Field1)
	}

	if testStruct.Field2 != 123 {
		t.Errorf("Hydrate() did not set Field2 correctly, expected 123, got %v", testStruct.Field2)
	}
}
