package hugolib

import (
	"errors"
	"fmt"
	"testing"
)

func TestFatalError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := &fatalErrorHandler{
		donec: make(chan bool),
	}

	err := errors.New("test error")
	handler.FatalError(err)

	if handler.getErr() != err {
		t.Errorf("Expected error %v, got %v", err, handler.getErr())
	}

	select {
	case <-handler.Done():
	default:
		t.Error("Expected Done() to be closed")
	}
}
