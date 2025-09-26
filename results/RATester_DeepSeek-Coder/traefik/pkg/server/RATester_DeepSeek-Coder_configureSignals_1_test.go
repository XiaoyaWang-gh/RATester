package server

import (
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestConfigureSignals_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Server{
		signals: make(chan os.Signal, 1),
	}

	s.configureSignals()

	syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)

	select {
	case sig := <-s.signals:
		if sig != syscall.SIGUSR1 {
			t.Errorf("Expected signal %v, got %v", syscall.SIGUSR1, sig)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for signal")
	}
}
