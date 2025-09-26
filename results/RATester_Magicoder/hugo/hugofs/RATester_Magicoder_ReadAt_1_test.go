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

	nop := &noOpRegularFileOps{}
	_, err := nop.ReadAt([]byte{}, 0)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err != errNoOp {
		t.Errorf("Expected error %v, got %v", errNoOp, err)
	}
}
