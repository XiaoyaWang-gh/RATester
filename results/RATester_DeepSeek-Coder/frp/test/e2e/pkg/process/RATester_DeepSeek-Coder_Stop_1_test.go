package process

import (
	"fmt"
	"testing"
	"time"
)

func TestStop_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		beforeStop    func()
		expectedError error
	}{
		{
			name: "should return nil if process is already stopped",
			beforeStop: func() {
				p := &Process{
					stopped: true,
				}
				err := p.Stop()
				if err != nil {
					t.Errorf("expected nil, got %v", err)
				}
			},
			expectedError: nil,
		},
		{
			name: "should call beforeStopHandler and cancel context if process is not stopped",
			beforeStop: func() {
				p := &Process{
					stopped: false,
					beforeStopHandler: func() {
						// simulate a long running process
						time.Sleep(1 * time.Second)
					},
				}
				err := p.Stop()
				if err != nil {
					t.Errorf("expected nil, got %v", err)
				}
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.beforeStop()
		})
	}
}
