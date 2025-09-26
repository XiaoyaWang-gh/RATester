package echo

import (
	"fmt"
	"testing"
)

func TestCreateFilesystem_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := createFilesystem()
	if fs.Filesystem == nil {
		t.Errorf("Expected Filesystem to be not nil")
	}
}
