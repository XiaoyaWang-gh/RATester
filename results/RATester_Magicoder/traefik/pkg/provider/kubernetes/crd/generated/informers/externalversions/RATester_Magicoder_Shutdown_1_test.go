package externalversions

import (
	"fmt"
	"testing"
)

func TestShutdown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new instance of sharedInformerFactory
	f := &sharedInformerFactory{
		shuttingDown: false,
	}

	// Call the method under test
	f.Shutdown()

	// Assert that the shuttingDown field is true
	if !f.shuttingDown {
		t.Errorf("Expected shuttingDown to be true, but it was false")
	}
}
