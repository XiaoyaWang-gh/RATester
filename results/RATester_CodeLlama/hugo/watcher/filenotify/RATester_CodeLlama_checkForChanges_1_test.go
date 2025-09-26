package filenotify

import (
	"fmt"
	"testing"
)

func TestCheckForChanges_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: add test cases.
}
