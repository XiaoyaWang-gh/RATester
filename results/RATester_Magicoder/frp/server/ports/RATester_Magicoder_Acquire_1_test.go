package ports

import (
	"fmt"
	"testing"
)

func TestAcquire_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &Manager{
		reservedPorts: make(map[string]*PortCtx),
		usedPorts:     make(map[int]*PortCtx),
		freePorts:     make(map[int]struct{}),
	}

	// Test case 1: Acquire a new port
	port, err := manager.Acquire("test1", 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if port == 0 {
		t.Errorf("Expected non-zero port, got 0")
	}

	// Test case 2: Acquire an existing port
	port, err = manager.Acquire("test1", port)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if port == 0 {
		t.Errorf("Expected non-zero port, got 0")
	}

	// Test case 3: Acquire a port that's already in use
	_, err = manager.Acquire("test2", port)
	if err != ErrPortAlreadyUsed {
		t.Errorf("Expected ErrPortAlreadyUsed, got %v", err)
	}

	// Test case 4: Acquire a port that's not allowed
	_, err = manager.Acquire("test3", 80)
	if err != ErrPortNotAllowed {
		t.Errorf("Expected ErrPortNotAllowed, got %v", err)
	}

	// Test case 5: Acquire a port that's not available
	manager.freePorts[80] = struct{}{}
	_, err = manager.Acquire("test4", 80)
	if err != ErrPortUnAvailable {
		t.Errorf("Expected ErrPortUnAvailable, got %v", err)
	}

	// Test case 6: Acquire a port that's not available after trying multiple times
	_, err = manager.Acquire("test5", 0)
	if err != ErrNoAvailablePort {
		t.Errorf("Expected ErrNoAvailablePort, got %v", err)
	}
}
