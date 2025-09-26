package utils

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestSelfDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := filepath.Dir(SelfPath())
	actual := SelfDir()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
