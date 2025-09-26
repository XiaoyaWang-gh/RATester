package grace

import (
	"fmt"
	"os"
	"testing"
)

func TestsignalHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		SignalHooks: make(map[int]map[os.Signal][]func()),
	}

	// Test case 1: Signal not set
	ppFlag := 1
	sig := os.Interrupt
	srv.signalHooks(ppFlag, sig)

	// Test case 2: Signal set
	srv.SignalHooks[ppFlag] = make(map[os.Signal][]func())
	srv.SignalHooks[ppFlag][sig] = []func(){func() {}}
	srv.signalHooks(ppFlag, sig)
}
