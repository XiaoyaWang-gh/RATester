package models

import (
	"fmt"
	"testing"
	"time"
)

func TestRawValue_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dateField := &DateField{}
	dateField.Set(time.Now())
	expected := dateField.Value()
	actual := dateField.RawValue()
	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
