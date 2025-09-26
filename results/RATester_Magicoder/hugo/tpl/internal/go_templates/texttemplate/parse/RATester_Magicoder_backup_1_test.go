package parse

import (
	"fmt"
	"testing"
)

func Testbackup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		peekCount: 0,
	}
	tree.backup()
	if tree.peekCount != 1 {
		t.Errorf("Expected peekCount to be 1, but got %d", tree.peekCount)
	}
}
