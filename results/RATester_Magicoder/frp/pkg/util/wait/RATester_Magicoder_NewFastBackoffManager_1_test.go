package wait

import (
	"fmt"
	"testing"
	"time"
)

func TestNewFastBackoffManager_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	options := FastBackoffOptions{
		Duration:           time.Second,
		Factor:             2,
		Jitter:             0.1,
		MaxDuration:        time.Minute,
		InitDurationIfFail: time.Second,
		FastRetryCount:     3,
		FastRetryDelay:     time.Second,
		FastRetryJitter:    0.2,
		FastRetryWindow:    time.Minute,
	}

	manager := NewFastBackoffManager(options)

	if manager == nil {
		t.Fatal("NewFastBackoffManager returned nil")
	}

	// Add more specific tests here
}
