package safe

import (
	"fmt"
	"testing"
	"time"
)

func TestGoWithRecover_1(t *testing.T) {
	testCases := []struct {
		name          string
		goroutine     func()
		customRecover func(err interface{})
		expectedPanic bool
	}{
		{
			name: "Test case 1: Normal execution",
			goroutine: func() {
				fmt.Println("Normal execution")
			},
			customRecover: func(err interface{}) {
				t.Errorf("Unexpected panic: %v", err)
			},
			expectedPanic: false,
		},
		{
			name: "Test case 2: Panic execution",
			goroutine: func() {
				panic("Panic execution")
			},
			customRecover: func(err interface{}) {
				if err != "Panic execution" {
					t.Errorf("Unexpected panic: %v", err)
				}
			},
			expectedPanic: true,
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
				if r := recover(); r != nil {
					if !tc.expectedPanic {
						t.Errorf("Unexpected panic: %v", r)
					}
				} else if tc.expectedPanic {
					t.Error("Expected panic but didn't occur")
				}
			}()

			GoWithRecover(tc.goroutine, tc.customRecover)
			time.Sleep(100 * time.Millisecond)
		})
	}
}
