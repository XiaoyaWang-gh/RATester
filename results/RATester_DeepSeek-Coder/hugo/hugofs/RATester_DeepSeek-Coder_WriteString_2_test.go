package hugofs

import (
	"fmt"
	"testing"
)

func TestWriteString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := &noOpRegularFileOps{}
	_, err := file.WriteString("test")
	if err != errNoOp {
		t.Errorf("Expected errNoOp, got %v", err)
	}
}
