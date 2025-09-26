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

	expression := ".*"
	d := &Decoder{}
	err := IncludeFields(expression)(d)
	if err != nil {
		t.Fatalf("IncludeFields(%q) failed: %s", expression, err)
	}
	if d.includeFieldsRe == nil {
		t.Fatalf("IncludeFields(%q) did not set includeFieldsRe", expression)
	}
}
