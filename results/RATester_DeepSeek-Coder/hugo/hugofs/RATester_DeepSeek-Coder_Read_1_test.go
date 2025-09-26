package hugofs

import (
	"fmt"
	"testing"
)

func TestRead_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := &noOpRegularFileOps{}
	_, err := file.Read([]byte{})
	if err != errNoOp {
		t.Errorf("Expected errNoOp, got %v", err)
	}
}
