package test

import (
	"fmt"
	"testing"
)

func TestSize_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := bindataFileInfo{
		name: "test.txt",
		size: 1024,
	}
	if fi.Size() != 1024 {
		t.Errorf("Size() = %d, want %d", fi.Size(), 1024)
	}
}
