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

	c := &RenameColumn{}
	dataType := "int"
	c.SetOldDataType(dataType)
	if c.OldDataType != dataType {
		t.Errorf("SetOldDataType failed")
	}
}
