package accesslog

import (
	"fmt"
	"testing"
)

func TestClose_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &Handler{
		logHandlerChan: make(chan handlerParams),
	}

	// Simulate a goroutine writing to the channel
	go func() {
		h.logHandlerChan <- handlerParams{}
	}()

	// Close the handler
	err := h.Close()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the channel is closed
	_, ok := <-h.logHandlerChan
	if ok {
		t.Error("Channel is not closed")
	}
}
