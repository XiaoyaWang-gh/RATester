package validation

import (
	"fmt"
	"testing"
)

func TestValid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string `valid:"Required"`
		Field2 int    `valid:"Required"`
	}

	v := Validation{}

	// Test case 1: Valid struct
	testStruct := TestStruct{Field1: "test", Field2: 1}
	b, err := v.Valid(testStruct)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !b {
		t.Errorf("Expected true, got false")
	}

	// Test case 2: Invalid struct
	testStruct = TestStruct{Field1: "", Field2: 0}
	b, err = v.Valid(testStruct)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if b {
		t.Errorf("Expected false, got true")
	}

	// Test case 3: Invalid pointer to struct
	testStructPtr := &TestStruct{Field1: "", Field2: 0}
	b, err = v.Valid(testStructPtr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if b {
		t.Errorf("Expected false, got true")
	}

	// Test case 4: Non-struct or non-pointer
	nonStruct := "test"
	b, err = v.Valid(nonStruct)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if b {
		t.Errorf("Expected false, got true")
	}
}
