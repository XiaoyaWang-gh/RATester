package test

import (
	"fmt"
	"os"
	"testing"
)

func TestRestoreAsset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dir := "testdata"
	name := "test.txt"
	err := RestoreAsset(dir, name)
	if err != nil {
		t.Errorf("RestoreAsset() error = %v", err)
		return
	}
	if _, err := os.Stat(_filePath(dir, name)); err != nil {
		t.Errorf("RestoreAsset() error = %v", err)
		return
	}
}
