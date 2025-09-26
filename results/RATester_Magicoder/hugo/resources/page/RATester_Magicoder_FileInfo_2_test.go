package page

import (
	"fmt"
	"testing"
)

func TestFileInfo_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	info := nop.FileInfo()

	if info == nil {
		t.Error("Expected FileInfo to not be nil")
	}

	if info.Name() != "" {
		t.Errorf("Expected Name to be empty, got %s", info.Name())
	}

	if info.IsDir() {
		t.Error("Expected IsDir to be false")
	}

	if info.Size() != 0 {
		t.Errorf("Expected Size to be 0, got %d", info.Size())
	}

	if info.Mode() != 0 {
		t.Errorf("Expected Mode to be 0, got %d", info.Mode())
	}

	if !info.ModTime().IsZero() {
		t.Errorf("Expected ModTime to be zero time, got %v", info.ModTime())
	}

	if info.Sys() != nil {
		t.Errorf("Expected Sys to be nil, got %v", info.Sys())
	}
}
