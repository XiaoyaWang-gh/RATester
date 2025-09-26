package traefik

import (
	"fmt"
	"testing"
	"time"
)

func TestThrottleDuration_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := Provider{}
	expected := time.Duration(0)
	result := provider.ThrottleDuration()
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
