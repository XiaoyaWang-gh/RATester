package safe

import (
	"fmt"
	"testing"
)

func TestGo_1(t *testing.T) {
	testCases := []struct {
		name      string
		goroutine func()
		wantPanic bool
	}{
		{
			name: "normal case",
			goroutine: func() {
				// do something
			},
			wantPanic: false,
		},
		{
			name: "panic case",
			goroutine: func() {
				panic("test panic")
			},
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
					t.Errorf("unexpected panic: %v, wantPanic: %v", r, tc.wantPanic)
				}
			}()

			Go(tc.goroutine)
		})
	}
}
