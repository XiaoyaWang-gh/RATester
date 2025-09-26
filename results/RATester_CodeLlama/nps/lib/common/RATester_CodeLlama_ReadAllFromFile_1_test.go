package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReadAllFromFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	filePath := "./testdata/test.txt"
	expected := []byte("Hello, World!")

	// Act
	actual, err := ReadAllFromFile(filePath)

	// Assert
	if err != nil {
		t.Errorf("ReadAllFromFile() error = %v", err)
		return
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ReadAllFromFile() = %v, want %v", actual, expected)
	}
}
