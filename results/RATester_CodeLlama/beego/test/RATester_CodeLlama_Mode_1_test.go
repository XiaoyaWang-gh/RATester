package test

import (
	"fmt"
	"testing"
)

func TestMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var fi bindataFileInfo
	if got := fi.Mode(); got != fi.mode {
		t.Errorf("bindataFileInfo.Mode() = %v, want %v", got, fi.mode)
	}
}
