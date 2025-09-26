package web

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestTestBeegoInit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ap := "/path/to/your/app"
	TestBeegoInit(ap)

	// Assertions
	if err := os.Chdir(ap); err != nil {
		t.Errorf("Failed to change directory: %v", err)
	}

	path := filepath.Join(ap, "conf", "app.conf")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("File does not exist: %v", path)
	}
}
