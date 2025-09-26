package models

import (
	"fmt"
	"testing"
)

func TestFieldType_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := BigIntegerField(0)
	expected := TypeBigIntegerField
	actual := e.FieldType()
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
