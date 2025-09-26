package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestReadAllFromFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filePath := "test.txt"
	content := []byte("Hello, world!\n")
	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filePath)

	data, err := ReadAllFromFile(filePath)
	if err != nil {
		t.Errorf("ReadAllFromFile() error = %v", err)
		return
	}

	if !reflect.DeepEqual(data, content) {
		t.Errorf("ReadAllFromFile() = %v, want %v", data, content)
	}
}
