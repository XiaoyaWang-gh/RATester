package server

import (
	"fmt"
	"testing"
	"time"
)

func TestdealClientFlow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	done := make(chan bool)
	go func() {
		dealClientFlow()
		done <- true
	}()

	select {
	case <-done:
		// Test passed
	case <-time.After(time.Second):
		t.Error("Test timed out")
	}
}
