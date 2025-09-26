package rate

import (
	"fmt"
	"testing"
	"time"
)

func TestSession_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Rate{
		bucketSize:        100,
		bucketSurplusSize: 50,
		bucketAddSize:     20,
		stopChan:          make(chan bool),
	}

	go r.session()

	time.Sleep(time.Second * 3)
	r.Stop()

	if r.NowRate != 30 {
		t.Errorf("Expected NowRate to be 30, got %d", r.NowRate)
	}
}
