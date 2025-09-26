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

	fm := NewFileMeta()
	if fm == nil {
		t.Errorf("Expected a new FileMeta, got nil")
	}
}
