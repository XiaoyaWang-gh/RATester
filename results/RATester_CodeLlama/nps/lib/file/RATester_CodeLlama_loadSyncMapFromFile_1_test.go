package file

import (
	"fmt"
	"testing"
)

func TestLoadSyncMapFromFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	filePath := "./testdata/test.txt"
	f := func(value string) {
		fmt.Println(value)
	}

	// Act
	loadSyncMapFromFile(filePath, f)

	// Assert
	// TODO
}
