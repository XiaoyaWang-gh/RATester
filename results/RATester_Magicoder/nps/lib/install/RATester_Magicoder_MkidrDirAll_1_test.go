package install

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestMkidrDirAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testPath := "/tmp/test"
	testDir := "testDir"
	testSubDir := "testSubDir"

	// Create test directory
	MkidrDirAll(testPath, testDir)

	// Check if the directory was created
	_, err := os.Stat(filepath.Join(testPath, testDir))
	if os.IsNotExist(err) {
		t.Errorf("Directory %s was not created", testDir)
	}

	// Create sub directory
	MkidrDirAll(testPath, testDir, testSubDir)

	// Check if the sub directory was created
	_, err = os.Stat(filepath.Join(testPath, testDir, testSubDir))
	if os.IsNotExist(err) {
		t.Errorf("Sub directory %s was not created", testSubDir)
	}
}
