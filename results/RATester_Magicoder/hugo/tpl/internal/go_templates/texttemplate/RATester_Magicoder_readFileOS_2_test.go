package template

import (
	"bytes"
	"fmt"
	"testing"
)

func TestreadFileOS_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := "test.txt"
	expectedName := "test.txt"
	expectedContent := []byte("This is a test file.")

	name, content, err := readFileOS(file)

	if err != nil {
		t.Errorf("readFileOS(%s) returned an error: %v", file, err)
	}

	if name != expectedName {
		t.Errorf("readFileOS(%s) returned name %s, expected %s", file, name, expectedName)
	}

	if !bytes.Equal(content, expectedContent) {
		t.Errorf("readFileOS(%s) returned content %s, expected %s", file, content, expectedContent)
	}
}
