package retry

import (
	"fmt"
	"testing"
)

func TestDisableRetries_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &responseWriter{}
	r.DisableRetries()
	if r.shouldRetry {
		t.Errorf("DisableRetries() should set shouldRetry to false")
	}
}
