package wrr

import (
	"fmt"
	"testing"
)

func TestRegisterStatusUpdater_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Balancer{
		wantsHealthCheck: true,
	}

	fn := func(up bool) {
		// do something
	}

	err := b.RegisterStatusUpdater(fn)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(b.updaters) != 1 {
		t.Errorf("Expected 1 updater, got %d", len(b.updaters))
	}

	b.wantsHealthCheck = false
	err = b.RegisterStatusUpdater(fn)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if len(b.updaters) != 1 {
		t.Errorf("Expected 1 updater, got %d", len(b.updaters))
	}
}
