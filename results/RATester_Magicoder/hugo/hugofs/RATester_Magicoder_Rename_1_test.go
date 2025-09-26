package hugofs

import (
	"fmt"
	"testing"
)

func TestRename_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := noOpFs{}
	err := fs.Rename("oldname", "newname")
	if err != errNoOp {
		t.Errorf("Expected errNoOp, got %v", err)
	}
}
