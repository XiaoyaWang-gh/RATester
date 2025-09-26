package migration

import (
	"fmt"
	"testing"
)

func TestSetUnsigned_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	column := &Column{}
	column.SetUnsigned(true)
	if column.Unsign != "UNSIGNED" {
		t.Errorf("Expected Unsign to be 'UNSIGNED', but got '%s'", column.Unsign)
	}
}
