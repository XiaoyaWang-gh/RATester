package hugofs

import (
	"fmt"
	"testing"
	"time"
)

func TestChtimes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := noOpFs{}
	name := "testfile"
	atime := time.Now()
	mtime := time.Now()

	err := fs.Chtimes(name, atime, mtime)
	if err != errNoOp {
		t.Errorf("Expected error %v, got %v", errNoOp, err)
	}
}
