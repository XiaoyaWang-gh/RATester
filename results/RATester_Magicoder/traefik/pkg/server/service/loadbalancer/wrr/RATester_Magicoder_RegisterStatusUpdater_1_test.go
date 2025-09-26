package wrr

import (
	"context"
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

	updaterCalled := false
	updater := func(up bool) {
		updaterCalled = true
	}

	err := b.RegisterStatusUpdater(updater)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(b.updaters) != 1 {
		t.Fatalf("Expected 1 updater, got %d", len(b.updaters))
	}

	if updaterCalled {
		t.Fatalf("Updater should not have been called yet")
	}

	b.SetStatus(context.Background(), "child", true)

	if !updaterCalled {
		t.Fatalf("Updater should have been called")
	}
}
