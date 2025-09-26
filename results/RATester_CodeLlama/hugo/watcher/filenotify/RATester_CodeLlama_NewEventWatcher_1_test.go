package filenotify

import (
	"fmt"
	"testing"
)

func TestNewEventWatcher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	watcher, err := NewEventWatcher()
	if err != nil {
		t.Fatal(err)
	}
	defer watcher.Close()

	if err := watcher.Add("testdata/test.txt"); err != nil {
		t.Fatal(err)
	}

	if err := watcher.Remove("testdata/test.txt"); err != nil {
		t.Fatal(err)
	}

	if err := watcher.Add("testdata/test.txt"); err != nil {
		t.Fatal(err)
	}

	if err := watcher.Close(); err != nil {
		t.Fatal(err)
	}
}
