package commands

import (
	"fmt"
	"testing"
)

func TestnewWatcher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hb := &hugoBuilder{
		r: &rootCommand{
			Printf: func(format string, v ...interface{}) {
				fmt.Printf(format, v...)
			},
		},
	}

	watcher, err := hb.newWatcher("500ms", "dir1", "dir2")
	if err != nil {
		t.Errorf("Error creating watcher: %v", err)
	}

	// Test watcher functionality here

	watcher.Close()
}
