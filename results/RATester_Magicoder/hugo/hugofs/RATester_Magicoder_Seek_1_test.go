package hugofs

import (
	"fmt"
	"testing"
)

func TestSeek_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := &noOpRegularFileOps{}
	_, err := file.Seek(0, 0)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err != errNoOp {
		t.Errorf("Expected error %v, got %v", errNoOp, err)
	}
}
