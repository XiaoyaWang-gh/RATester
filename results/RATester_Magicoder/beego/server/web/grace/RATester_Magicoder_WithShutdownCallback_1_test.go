package grace

import (
	"fmt"
	"testing"
)

func TestWithShutdownCallback_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{}
	shutdownCallback := func() {
		// Test logic here
	}
	WithShutdownCallback(shutdownCallback)(srv)

	if len(srv.shutdownCallbacks) != 1 {
		t.Errorf("Expected 1 shutdown callback, got %d", len(srv.shutdownCallbacks))
	}

	srv.shutdownCallbacks[0]()

	// Add assertions here to verify the behavior of the shutdownCallback
}
