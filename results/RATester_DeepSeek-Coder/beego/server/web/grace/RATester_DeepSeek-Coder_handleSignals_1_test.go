package grace

import (
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestHandleSignals_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		sigChan: make(chan os.Signal, 1),
	}

	go srv.handleSignals()

	testSignals := []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}
	for _, sig := range testSignals {
		srv.sigChan <- sig
		time.Sleep(10 * time.Millisecond)
	}

	srv.shutdown()
}
