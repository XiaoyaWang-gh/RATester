package test

import (
	"fmt"
	"testing"
)

func TestRestoreAssets_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dir := "testdata"
	name := "testdata/test.txt"
	err := RestoreAssets(dir, name)
	if err != nil {
		t.Errorf("RestoreAssets() error = %v", err)
		return
	}
}
