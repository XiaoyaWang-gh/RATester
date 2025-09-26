package filenotify

import (
	"fmt"
	"testing"
)

func TestWatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup
	// TODO: test
	// TODO: teardown
}
