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
		bucketSurplusSize: 100,
		bucketAddSize:     10,
		stopChan:          make(chan bool),
	}

	r.Start()

	time.Sleep(time.Second)

	r.Stop()

	if r.NowRate != 0 {
		t.Errorf("Expected NowRate to be 0, but got %d", r.NowRate)
	}
}
