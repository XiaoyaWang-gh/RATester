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

	expression := "test"
	d := &Decoder{}
	err := ExcludeFields(expression)(d)
	if err != nil {
		t.Errorf("ExcludeFields failed: %v", err)
	}
	if d.excludeFieldsrRe == nil {
		t.Errorf("ExcludeFields failed: d.excludeFieldsrRe is nil")
	}
}
