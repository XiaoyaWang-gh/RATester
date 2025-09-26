package models

import (
	"fmt"
	"testing"
)

func TestFieldType_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dt := DateTimeField{}
	expected := TypeDateTimeField
	actual := dt.FieldType()
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
