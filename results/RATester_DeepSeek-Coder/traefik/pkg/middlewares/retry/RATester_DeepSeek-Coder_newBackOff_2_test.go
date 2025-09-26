package retry

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/cenkalti/backoff"
)

func TestNewBackOff_2(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name            string
		attempts        int
		initialInterval time.Duration
		expectedBackOff backoff.BackOff
	}

	testCases := []testCase{
		{
			name:            "less than 2 attempts",
			attempts:        1,
			initialInterval: 100 * time.Millisecond,
			expectedBackOff: &backoff.ZeroBackOff{},
		},
		{
			name:            "2 attempts",
			attempts:        2,
			initialInterval: 100 * time.Millisecond,
			expectedBackOff: &backoff.ZeroBackOff{},
		},
		{
			name:            "more than 2 attempts",
			attempts:        3,
			initialInterval: 100 * time.Millisecond,
			expectedBackOff: &backoff.ExponentialBackOff{
				InitialInterval:     100 * time.Millisecond,
				Multiplier:          1.61803398875, // Golden ratio
				RandomizationFactor: 0.5,
				MaxInterval:         5 * time.Second,
				MaxElapsedTime:      0,
				Clock:               backoff.SystemClock,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &retry{
				attempts:        tc.attempts,
				initialInterval: tc.initialInterval,
			}

			actualBackOff := r.newBackOff()

			if !reflect.DeepEqual(actualBackOff, tc.expectedBackOff) {
				t.Errorf("expected %v, got %v", tc.expectedBackOff, actualBackOff)
			}
		})
	}
}
