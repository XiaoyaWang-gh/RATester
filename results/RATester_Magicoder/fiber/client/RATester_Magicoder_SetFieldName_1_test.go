package client

import (
	"fmt"
	"testing"
)

func TestSetFieldName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := &File{}
	file.SetFieldName("test")

	if file.fieldName != "test" {
		t.Errorf("Expected fieldName to be 'test', but got '%s'", file.fieldName)
	}
}
