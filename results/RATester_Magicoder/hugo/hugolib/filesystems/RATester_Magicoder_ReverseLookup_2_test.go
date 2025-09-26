package filesystems

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestReverseLookup_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	d := &SourceFilesystem{
		Name: "test",
		Fs:   fs,
	}

	testFile := "test.txt"
	testContent := "test content"

	err := afero.WriteFile(fs, testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	_, err = d.ReverseLookup(testFile, true)
	if err != nil {
		t.Fatalf("ReverseLookup failed: %v", err)
	}

	// Test with non-existent file
	_, err = d.ReverseLookup("non-existent.txt", true)
	if err == nil {
		t.Fatalf("Expected error for non-existent file, but got nil")
	}
}
