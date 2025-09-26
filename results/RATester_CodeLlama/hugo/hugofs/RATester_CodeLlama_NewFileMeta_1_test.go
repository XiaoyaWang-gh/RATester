package hugofs

import (
	"fmt"
	"testing"
)

func TestNewFileMeta_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fileMeta := NewFileMeta()
	if fileMeta == nil {
		t.Errorf("NewFileMeta() = nil")
	}
}
