package retry

import (
	"fmt"
	"testing"
)

func TestShouldRetry_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &responseWriter{}
	if r.ShouldRetry() {
		t.Errorf("ShouldRetry() = true, want false")
	}
}
