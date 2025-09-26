package migration

import (
	"fmt"
	"testing"
)

func TestSetOldNullable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rc := &RenameColumn{}

	// Test with null = true
	rc.SetOldNullable(true)
	if rc.OldNull != "" {
		t.Errorf("Expected OldNull to be empty, but got %s", rc.OldNull)
	}

	// Test with null = false
	rc.SetOldNullable(false)
	if rc.OldNull != "NOT NULL" {
		t.Errorf("Expected OldNull to be 'NOT NULL', but got %s", rc.OldNull)
	}
}
