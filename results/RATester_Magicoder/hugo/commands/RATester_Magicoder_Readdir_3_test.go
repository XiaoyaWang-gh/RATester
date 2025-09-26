package commands

import (
	"fmt"
	"testing"
)

func TestReaddir_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := noDirFile{}
	_, err := n.Readdir(0)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
