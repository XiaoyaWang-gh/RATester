package grace

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

func TestRegisterSignalHook_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		SignalHooks: make(map[int]map[os.Signal][]func()),
	}

	// Test with valid ppFlag and valid signal
	err := srv.RegisterSignalHook(PreSignal, os.Interrupt, func() {})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test with invalid ppFlag
	err = srv.RegisterSignalHook(100, os.Interrupt, func() {})
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test with invalid signal
	err = srv.RegisterSignalHook(PreSignal, syscall.SIGKILL, func() {})
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
