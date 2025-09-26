package failover

import (
	"fmt"
	"testing"
)

func TestRegisterStatusUpdater_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &Failover{wantsHealthCheck: true}

	err := f.RegisterStatusUpdater(func(up bool) {})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(f.updaters) != 1 {
		t.Errorf("Expected 1 updater, got %d", len(f.updaters))
	}

	f.wantsHealthCheck = false
	err = f.RegisterStatusUpdater(func(up bool) {})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if len(f.updaters) != 1 {
		t.Errorf("Expected 1 updater, got %d", len(f.updaters))
	}
}
