package models

import (
	"fmt"
	"testing"
)

func TestFieldType_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	j := JSONField("test")
	expected := TypeJSONField
	actual := j.FieldType()
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
