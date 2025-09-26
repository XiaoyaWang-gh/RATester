package test

import (
	"fmt"
	"testing"
)

func TestIsDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := bindataFileInfo{
		name: "test",
	}
	if fi.IsDir() {
		t.Errorf("IsDir() = %v, want %v", fi.IsDir(), false)
	}
}
