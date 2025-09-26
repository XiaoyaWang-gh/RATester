package filenotify

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSendErr_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		setup   func(t *testing.T) *filePoller
		sendErr error
		wantErr error
	}{
		{
			name: "sends error when filePoller is not closed",
			setup: func(t *testing.T) *filePoller {
				fp := &filePoller{
					done:   make(chan struct{}),
					errors: make(chan error, 1),
				}
				return fp
			},
			sendErr: fmt.Errorf("test error"),
			wantErr: nil,
		},
		{
			name: "returns error when filePoller is closed",
			setup: func(t *testing.T) *filePoller {
				fp := &filePoller{
					done:   make(chan struct{}),
					errors: make(chan error, 1),
				}
				close(fp.done)
				return fp
			},
			sendErr: fmt.Errorf("test error"),
			wantErr: fmt.Errorf("closed"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fp := tc.setup(t)
			err := fp.sendErr(tc.sendErr)
			if !reflect.DeepEqual(err, tc.wantErr) {
				t.Errorf("Expected error %v, got %v", tc.wantErr, err)
			}
		})
	}
}
