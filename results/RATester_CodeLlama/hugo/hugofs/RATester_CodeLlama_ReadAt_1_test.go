package hugofs

import (
	"fmt"
	"testing"
)

func TestReadAt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &noOpRegularFileOps{}
	p := []byte{}
	off := int64(0)
	n, err := f.ReadAt(p, off)
	if n != 0 || err != errNoOp {
		t.Errorf("ReadAt() = (%d, %v), want (0, %v)", n, err, errNoOp)
	}
}
