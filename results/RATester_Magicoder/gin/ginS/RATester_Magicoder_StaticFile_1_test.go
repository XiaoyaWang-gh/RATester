package ginS

import (
	"fmt"
	"testing"
)

func TestStaticFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	relativePath := "/test"
	filepath := "/path/to/file"

	r := StaticFile(relativePath, filepath)

	if r == nil {
		t.Error("Expected a non-nil result, but got nil")
	}
}
