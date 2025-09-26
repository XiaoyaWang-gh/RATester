package server

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestListenSignals_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := &Server{
		signals: make(chan os.Signal, 1),
	}

	go s.listenSignals(ctx)

	s.signals <- syscall.SIGUSR1

	time.Sleep(100 * time.Millisecond)

	select {
	case <-s.signals:
		t.Error("Signal was not expected")
	default:
	}
}
