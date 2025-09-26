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

	ap := "./testdata"
	TestBeegoInit(ap)

	expectedPath := filepath.Join(ap, "conf", "app.conf")
	_, err := os.Stat(expectedPath)
	if err != nil {
		t.Errorf("Expected file %s not found", expectedPath)
	}

	expectedDir := filepath.Dir(expectedPath)
	_, err = os.Stat(expectedDir)
	if err != nil {
		t.Errorf("Expected directory %s not found", expectedDir)
	}
}
