package migration

import (
	"fmt"
	"testing"
)

func TestSetDataType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	column := &Column{}
	expectedDataType := "varchar"
	column.SetDataType(expectedDataType)
	if column.DataType != expectedDataType {
		t.Errorf("Expected DataType to be %s, but got %s", expectedDataType, column.DataType)
	}
}
