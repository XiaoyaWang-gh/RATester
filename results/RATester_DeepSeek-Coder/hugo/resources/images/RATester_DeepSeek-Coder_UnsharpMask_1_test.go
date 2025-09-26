package images

import (
	"fmt"
	"testing"
)

func TestUnsharpMask_1(t *testing.T) {
	t.Parallel()

	filters := &Filters{}

	testCases := []struct {
		name      string
		sigma     any
		amount    any
		threshold any
		wantPanic bool
	}{
		{
			name:      "valid inputs",
			sigma:     2.0,
			amount:    0.5,
			threshold: 0.1,
			wantPanic: false,
		},
		{
			name:      "invalid sigma",
			sigma:     "2.0",
			amount:    0.5,
			threshold: 0.1,
			wantPanic: true,
		},
		{
			name:      "invalid amount",
			sigma:     2.0,
			amount:    "0.5",
			threshold: 0.1,
			wantPanic: true,
		},
		{
			name:      "invalid threshold",
			sigma:     2.0,
			amount:    0.5,
			threshold: "0.1",
			wantPanic: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); (r != nil) != tc.wantPanic {
					t.Errorf("unexpected panic: %v", r)
				}
			}()

			filters.UnsharpMask(tc.sigma, tc.amount, tc.threshold)
		})
	}
}
