package hugofs

import (
	"fmt"
	"testing"
)

func TestTruncate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := &noOpRegularFileOps{}
	err := file.Truncate(100)
	if err != errNoOp {
		t.Errorf("Expected errNoOp, got %v", err)
	}
}
