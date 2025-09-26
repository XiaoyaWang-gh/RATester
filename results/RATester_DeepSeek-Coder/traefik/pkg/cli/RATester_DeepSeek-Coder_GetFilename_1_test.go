package cli

import (
	"fmt"
	"testing"
)

func TestGetFilename_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	loader := &FileLoader{
		ConfigFileFlag: "test.conf",
		filename:       "test.conf",
	}

	filename := loader.GetFilename()
	if filename != "test.conf" {
		t.Errorf("Expected filename to be 'test.conf', got '%s'", filename)
	}
}
