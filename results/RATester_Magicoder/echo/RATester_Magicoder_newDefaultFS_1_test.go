package echo

import (
	"fmt"
	"testing"
)

func TestNewDefaultFS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := newDefaultFS()
	if fs == nil {
		t.Error("newDefaultFS() should not return nil")
	}
	if fs.prefix != "." {
		t.Errorf("newDefaultFS().prefix should be '.' but got '%s'", fs.prefix)
	}
	if fs.fs != nil {
		t.Error("newDefaultFS().fs should be nil")
	}
}
