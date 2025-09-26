package grace

import (
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"
)

func TesthandleSignals_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		SignalHooks: make(map[int]map[os.Signal][]func()),
		sigChan:     make(chan os.Signal),
	}

	go srv.handleSignals()

	testSignal := syscall.SIGHUP
	srv.sigChan <- testSignal

	time.Sleep(1 * time.Second)

	if len(srv.SignalHooks[PreSignal][testSignal]) == 0 || len(srv.SignalHooks[PostSignal][testSignal]) == 0 {
		t.Errorf("Signal hooks not registered for signal %v", testSignal)
	}
}
