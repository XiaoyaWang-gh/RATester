package rate

import (
	"fmt"
	"testing"
	"time"
)

func TestStart_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Rate{
		bucketSize:        100,
		bucketSurplusSize: 0,
		bucketAddSize:     10,
		stopChan:          make(chan bool),
	}

	r.Start()

	time.Sleep(time.Second)

	r.Stop()

	if r.NowRate != 10 {
		t.Errorf("Expected rate to be 10, got %d", r.NowRate)
	}
}
