package common

import (
	"fmt"
	"testing"
)

func TestGetTmpPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := GetTmpPath()
	if path != "/tmp" {
		t.Errorf("GetTmpPath() = %v, want %v", path, "/tmp")
	}
}
