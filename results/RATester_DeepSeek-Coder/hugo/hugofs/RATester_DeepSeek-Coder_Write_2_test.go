package hugofs

import (
	"fmt"
	"testing"
)

func TestWrite_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := &noOpRegularFileOps{}
	_, err := file.Write([]byte("test"))
	if err != errNoOp {
		t.Errorf("Expected errNoOp, got %v", err)
	}
}
