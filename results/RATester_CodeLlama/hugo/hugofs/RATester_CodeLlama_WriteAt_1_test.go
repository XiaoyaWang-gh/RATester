package hugofs

import (
	"fmt"
	"testing"
)

func TestWriteAt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &noOpRegularFileOps{}
	p := []byte("test")
	off := int64(0)
	n, err := f.WriteAt(p, off)
	if n != 0 {
		t.Errorf("WriteAt() = %v, want %v", n, 0)
	}
	if err != errNoOp {
		t.Errorf("WriteAt() = %v, want %v", err, errNoOp)
	}
}
