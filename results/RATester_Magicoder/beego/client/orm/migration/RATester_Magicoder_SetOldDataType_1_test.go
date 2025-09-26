package migration

import (
	"fmt"
	"testing"
)

func TestSetOldDataType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rc := &RenameColumn{}
	dataType := "varchar"
	rc.SetOldDataType(dataType)
	if rc.OldDataType != dataType {
		t.Errorf("Expected OldDataType to be %s, but got %s", dataType, rc.OldDataType)
	}
}
