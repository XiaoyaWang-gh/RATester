package common

import (
	"fmt"
	"testing"
)

func TestGetAppPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := GetAppPath()
	if path == "" {
		t.Errorf("GetAppPath() = %v, want non-empty string", path)
	}
}
