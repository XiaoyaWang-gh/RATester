package acme

import (
	"fmt"
	"testing"
	"time"
)

func TestThrottleDuration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		Configuration: &Configuration{
			CertificatesDuration: 24,
		},
	}

	expectedDuration := 24 * time.Hour
	actualDuration := provider.ThrottleDuration()

	if actualDuration != expectedDuration {
		t.Errorf("Expected ThrottleDuration to be %v, but got %v", expectedDuration, actualDuration)
	}
}
