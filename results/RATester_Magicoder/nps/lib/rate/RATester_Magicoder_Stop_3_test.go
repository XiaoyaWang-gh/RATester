package rate

import (
	"fmt"
	"testing"
	"time"
)

func TestStop_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rate := &Rate{
		bucketSize:        100,
		bucketSurplusSize: 100,
		bucketAddSize:     10,
		stopChan:          make(chan bool),
	}

	go rate.Start()

	time.Sleep(time.Second)

	rate.Stop()

	select {
	case <-rate.stopChan:
	default:
		t.Error("Stop function is not working as expected")
	}
}
