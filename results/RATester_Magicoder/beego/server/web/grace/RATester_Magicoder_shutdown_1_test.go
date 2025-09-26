package grace

import (
	"fmt"
	"testing"
)

func Testshutdown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		state: StateRunning,
	}

	srv.shutdown()

	if srv.state != StateShuttingDown {
		t.Errorf("Expected state to be StateShuttingDown, but got %v", srv.state)
	}
}
