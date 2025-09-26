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

	statsdTicker = time.NewTicker(time.Second)
	StopStatsd()
	if statsdTicker != nil {
		t.Errorf("Expected statsdTicker to be nil, but got %v", statsdTicker)
	}
}
