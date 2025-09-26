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
		t.Errorf("NewEventWatcher() returned an error: %v", err)
	}

	if watcher == nil {
		t.Error("NewEventWatcher() returned a nil watcher")
	}

	events := watcher.Events()
	if events == nil {
		t.Error("NewEventWatcher() returned a nil events channel")
	}

	errors := watcher.Errors()
	if errors == nil {
		t.Error("NewEventWatcher() returned a nil errors channel")
	}

	err = watcher.Add("test.txt")
	if err != nil {
		t.Errorf("watcher.Add() returned an error: %v", err)
	}

	err = watcher.Remove("test.txt")
	if err != nil {
		t.Errorf("watcher.Remove() returned an error: %v", err)
	}

	err = watcher.Close()
	if err != nil {
		t.Errorf("watcher.Close() returned an error: %v", err)
	}
}
