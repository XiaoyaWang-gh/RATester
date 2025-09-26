package hugolib

import (
	"fmt"
	"testing"
	"time"
)

func TestDone_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &fatalErrorHandler{
		donec: make(chan bool),
	}

	go func() {
		time.Sleep(1 * time.Second)
		f.donec <- true
	}()

	select {
	case <-f.Done():
		// Test passed
		return
	case <-time.After(2 * time.Second):
		// Test failed
		t.Error("Test timed out")
	}
}
