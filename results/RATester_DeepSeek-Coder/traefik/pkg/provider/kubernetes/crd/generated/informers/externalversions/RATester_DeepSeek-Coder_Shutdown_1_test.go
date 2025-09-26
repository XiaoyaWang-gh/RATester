package externalversions

import (
	"fmt"
	reflect "reflect"
	sync "sync"
	"testing"
)

func TestShutdown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	factory := &sharedInformerFactory{
		lock:             sync.Mutex{},
		shuttingDown:     false,
		wg:               sync.WaitGroup{},
		startedInformers: make(map[reflect.Type]bool),
	}

	factory.Shutdown()

	if !factory.shuttingDown {
		t.Errorf("Expected shuttingDown to be true after calling Shutdown")
	}
}
