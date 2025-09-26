package rate

import (
	"fmt"
	"testing"
)

func TestStop_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rate := &Rate{
		bucketSize:        100,
		bucketSurplusSize: 50,
		bucketAddSize:     10,
		stopChan:          make(chan bool),
		NowRate:           10,
	}

	go rate.Stop()

	select {
	case <-rate.stopChan:
	default:
		t.Errorf("Stop() did not close the stopChan")
	}
}
