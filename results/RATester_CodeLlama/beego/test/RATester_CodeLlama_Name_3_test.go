package test

import (
	"fmt"
	"testing"
)

func TestName_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := bindataFileInfo{name: "test"}
	if fi.Name() != "test" {
		t.Errorf("expected %s, got %s", "test", fi.Name())
	}
}
