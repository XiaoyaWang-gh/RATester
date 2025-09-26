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

	failover := &Failover{wantsHealthCheck: true}

	err := failover.RegisterStatusUpdater(func(up bool) {})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(failover.updaters) != 1 {
		t.Errorf("Expected 1 updater, got %d", len(failover.updaters))
	}

	failover.wantsHealthCheck = false
	err = failover.RegisterStatusUpdater(func(up bool) {})
	if err == nil {
		t.Error("Expected error, got nil")
	}

	if len(failover.updaters) != 1 {
		t.Errorf("Expected 1 updater, got %d", len(failover.updaters))
	}
}
