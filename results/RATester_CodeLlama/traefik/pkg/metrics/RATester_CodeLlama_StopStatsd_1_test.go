package metrics

import (
	"fmt"
	"testing"
	"time"
)

func TestStopStatsd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	statsdTicker := time.NewTicker(time.Second)
	// Act
	StopStatsd()
	// Assert
	if statsdTicker != nil {
		t.Errorf("statsdTicker should be nil")
	}
}
