package config

import (
	"fmt"
	"testing"
)

func TestSaveConfigFile_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &fakeConfigContainer{
		data: map[string]string{
			"key": "value",
		},
	}

	err := fc.SaveConfigFile("test.txt")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
