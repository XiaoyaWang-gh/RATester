package hugofs

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestChmod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	mode := fs.FileMode(0755)
	fs := noOpFs{}
	err := fs.Chmod(name, mode)
	if err != errNoOp {
		t.Errorf("Chmod() = %v, want %v", err, errNoOp)
	}
}
