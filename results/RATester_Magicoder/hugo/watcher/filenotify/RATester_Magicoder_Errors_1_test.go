package filenotify

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &filePoller{
		errors: make(chan error),
	}

	go func() {
		w.errors <- errors.New("test error")
	}()

	err := <-w.Errors()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
