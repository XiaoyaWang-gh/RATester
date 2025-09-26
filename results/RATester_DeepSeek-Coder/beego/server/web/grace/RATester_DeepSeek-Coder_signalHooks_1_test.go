package grace

import (
	"fmt"
	"os"
	"testing"
)

func TestSignalHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		SignalHooks: make(map[int]map[os.Signal][]func()),
	}

	ppFlag := 1
	sig := os.Interrupt
	f := func() {
		fmt.Println("Signal received")
	}

	srv.SignalHooks[ppFlag] = make(map[os.Signal][]func())
	srv.SignalHooks[ppFlag][sig] = append(srv.SignalHooks[ppFlag][sig], f)

	srv.signalHooks(ppFlag, sig)

	// Output: Signal received
}
