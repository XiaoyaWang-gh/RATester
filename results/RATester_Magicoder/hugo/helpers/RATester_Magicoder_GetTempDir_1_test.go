package helpers

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetTempDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	subPath := "test"
	tempDir := GetTempDir(subPath, fs)

	if tempDir == "" {
		t.Error("Expected a non-empty temp directory, got an empty string")
	}

	_, err := fs.Stat(tempDir)
	if err != nil {
		t.Errorf("Expected temp directory to exist, got error: %v", err)
	}
}
