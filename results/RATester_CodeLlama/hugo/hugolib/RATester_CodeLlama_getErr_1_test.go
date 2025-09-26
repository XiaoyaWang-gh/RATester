package hugolib

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &fatalErrorHandler{}
	f.mu.Lock()
	f.err = errors.New("test error")
	f.mu.Unlock()
	if err := f.getErr(); err == nil {
		t.Errorf("getErr() = nil, want %v", f.err)
	}
}
