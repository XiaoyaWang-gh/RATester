package wait

import (
	"fmt"
	"testing"
	"time"
)

func TestNewFastBackoffManager_1(t *testing.T) {
	t.Run("test NewFastBackoffManager", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		options := FastBackoffOptions{
			Duration:           time.Second,
			Factor:             1.5,
			Jitter:             0.5,
			MaxDuration:        time.Minute,
			InitDurationIfFail: time.Second,
			FastRetryCount:     3,
			FastRetryDelay:     time.Second,
			FastRetryJitter:    0.5,
			FastRetryWindow:    time.Minute,
		}
		manager := NewFastBackoffManager(options)
		if manager == nil {
			t.Error("NewFastBackoffManager failed")
		}
	})
}
