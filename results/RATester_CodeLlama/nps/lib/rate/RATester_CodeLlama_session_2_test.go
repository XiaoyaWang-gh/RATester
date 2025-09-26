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

	s := &Rate{
		bucketSize:        100,
		bucketSurplusSize: 10,
		bucketAddSize:     10,
		stopChan:          make(chan bool),
	}
	s.Start()
	time.Sleep(time.Second * 2)
	s.Stop()
	if s.NowRate != 10 {
		t.Errorf("NowRate is %d, want %d", s.NowRate, 10)
	}
}
