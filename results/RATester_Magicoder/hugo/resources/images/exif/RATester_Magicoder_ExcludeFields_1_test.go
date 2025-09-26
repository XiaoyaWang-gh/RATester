package exif

import (
	"fmt"
	"testing"
)

func TestExcludeFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Decoder{}
	expression := "test"
	ExcludeFields(expression)(d)

	if d.excludeFieldsrRe == nil {
		t.Error("Expected ExcludeFieldsrRe to be set, but it was nil")
	}
}
