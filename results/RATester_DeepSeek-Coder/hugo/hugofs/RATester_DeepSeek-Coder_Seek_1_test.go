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
	offset := int64(0)
	whence := 0
	_, err := file.Seek(offset, whence)
	if err != errNoOp {
		t.Errorf("Expected errNoOp, got %v", err)
	}
}
