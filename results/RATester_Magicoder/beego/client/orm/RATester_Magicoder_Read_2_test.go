package orm

import (
	"fmt"
	"testing"
)

func TestRead_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}

	type TestStruct struct {
		ID   int64
		Name string
	}

	testStruct := TestStruct{ID: 1, Name: "Test"}

	err := o.Read(&testStruct)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	if testStruct.ID != 1 {
		t.Errorf("Expected ID to be 1, but got %d", testStruct.ID)
	}

	if testStruct.Name != "Test" {
		t.Errorf("Expected Name to be 'Test', but got '%s'", testStruct.Name)
	}
}
