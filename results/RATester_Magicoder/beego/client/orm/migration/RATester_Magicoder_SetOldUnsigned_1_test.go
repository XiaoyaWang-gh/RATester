package migration

import (
	"fmt"
	"testing"
)

func TestSetOldUnsigned_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	renameColumn := &RenameColumn{}

	renameColumn.SetOldUnsigned(true)

	if renameColumn.OldUnsign != "UNSIGNED" {
		t.Errorf("Expected OldUnsign to be 'UNSIGNED', but got '%s'", renameColumn.OldUnsign)
	}

	renameColumn.SetOldUnsigned(false)

	if renameColumn.OldUnsign != "" {
		t.Errorf("Expected OldUnsign to be '', but got '%s'", renameColumn.OldUnsign)
	}
}
