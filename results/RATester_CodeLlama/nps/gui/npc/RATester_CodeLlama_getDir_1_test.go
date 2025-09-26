package main

import (
	"fmt"
	"testing"
)

func TestGetDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dir, err := getDir()
	if err != nil {
		t.Errorf("getDir() error = %v", err)
		return
	}
	if dir != "/data/data/org.nps.client/files" {
		t.Errorf("getDir() = %v, want %v", dir, "/data/data/org.nps.client/files")
	}
}
