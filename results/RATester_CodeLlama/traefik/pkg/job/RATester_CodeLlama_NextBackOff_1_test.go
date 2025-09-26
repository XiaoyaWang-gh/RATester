package job

import (
	"fmt"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func TestNextBackOff_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BackOff{
		ExponentialBackOff: &backoff.ExponentialBackOff{
			InitialInterval:     100 * time.Millisecond,
			RandomizationFactor: 0.5,
			Multiplier:          1.5,
			MaxInterval:         10 * time.Second,
			MaxElapsedTime:      10 * time.Second,
			Stop:                10 * time.Second,
			Clock:               backoff.SystemClock,
		},
		MinJobInterval: 10 * time.Second,
	}

	b.Reset()

	for i := 0; i < 10; i++ {
		d := b.NextBackOff()
		if d > b.MinJobInterval {
			t.Errorf("NextBackOff() = %v, want <= %v", d, b.MinJobInterval)
		}
	}
}
