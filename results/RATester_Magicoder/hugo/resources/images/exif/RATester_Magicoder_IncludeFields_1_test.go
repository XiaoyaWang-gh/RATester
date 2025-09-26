package exif

import (
	"fmt"
	"testing"
)

func TestIncludeFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Decoder{}
	expression := "test"
	err := IncludeFields(expression)(d)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if d.includeFieldsRe == nil {
		t.Errorf("Expected includeFieldsRe to be set, but it was nil")
	}
}
