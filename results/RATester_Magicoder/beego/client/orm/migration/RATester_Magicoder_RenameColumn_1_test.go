package migration

import (
	"fmt"
	"testing"
)

func TestRenameColumn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	from := "old_column"
	to := "new_column"
	rename := m.RenameColumn(from, to)

	if rename.OldName != from || rename.NewName != to {
		t.Errorf("Expected OldName to be %s and NewName to be %s, but got %s and %s", from, to, rename.OldName, rename.NewName)
	}

	if len(m.Renames) != 1 || m.Renames[0] != rename {
		t.Errorf("Expected Rename to be added to Renames, but it was not")
	}
}
