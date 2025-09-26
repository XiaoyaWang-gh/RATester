package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadAllFromFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filePath := "test.txt"
	content := []byte("This is a test file.")
	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filePath)

	data, err := ReadAllFromFile(filePath)
	if err != nil {
		t.Errorf("ReadAllFromFile returned an error: %v", err)
	}
	if !bytes.Equal(data, content) {
		t.Errorf("ReadAllFromFile returned incorrect data. Expected: %s, Got: %s", content, data)
	}
}
