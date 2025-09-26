package grace

import (
	"fmt"
	"os"
	"sync"
	"syscall"
	"testing"
)

func TestInit_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	isChild = false
	socketOrder = ""
	regLock = &sync.Mutex{}
	runningServers = make(map[string]*Server)
	runningServersOrder = []string{}
	socketPtrOffsetMap = make(map[string]uint)

	hookableSignals = []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
	}

	if isChild != false {
		t.Errorf("Expected isChild to be false, got %v", isChild)
	}

	if socketOrder != "" {
		t.Errorf("Expected socketOrder to be empty, got %v", socketOrder)
	}

	if regLock == nil {
		t.Errorf("Expected regLock to not be nil")
	}

	if len(runningServers) != 0 {
		t.Errorf("Expected runningServers to be empty, got %v", runningServers)
	}

	if len(runningServersOrder) != 0 {
		t.Errorf("Expected runningServersOrder to be empty, got %v", runningServersOrder)
	}

	if len(socketPtrOffsetMap) != 0 {
		t.Errorf("Expected socketPtrOffsetMap to be empty, got %v", socketPtrOffsetMap)
	}

	if len(hookableSignals) != 3 {
		t.Errorf("Expected hookableSignals to have 3 elements, got %v", len(hookableSignals))
	}
}
