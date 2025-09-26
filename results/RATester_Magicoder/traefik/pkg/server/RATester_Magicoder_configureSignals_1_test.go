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
	case <-s.signals:
		// Test passed
	case <-time.After(time.Second):
		t.Error("Signal not received")
	}
}
