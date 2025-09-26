package hugofs

import (
	"fmt"
	"testing"
)

func TestSync_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &noOpRegularFileOps{}
	err := f.Sync()
	if err != errNoOp {
		t.Errorf("Sync() = %v, want %v", err, errNoOp)
	}
}
