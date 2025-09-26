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

	t.Parallel()

	srv := &Server{}
	shutdownCallback := func() {
		t.Log("Shutdown callback called")
	}

	WithShutdownCallback(shutdownCallback)(srv)

	if len(srv.shutdownCallbacks) != 1 {
		t.Errorf("Expected 1 shutdown callback, got %d", len(srv.shutdownCallbacks))
	}

	srv.shutdownCallbacks[0]()
}
