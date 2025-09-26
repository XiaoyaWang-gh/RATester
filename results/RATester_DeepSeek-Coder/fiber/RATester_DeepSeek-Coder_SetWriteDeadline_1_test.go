package fiber

import (
	"fmt"
	"testing"
	"time"
)

func TestSetWriteDeadline_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		time    time.Time
		wantErr bool
	}{
		{
			name:    "success",
			time:    time.Now(),
			wantErr: false,
		},
		{
			name:    "failure",
			time:    time.Time{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			c := &testConn{}
			err := c.SetWriteDeadline(tc.time)
			if (err != nil) != tc.wantErr {
				t.Errorf("SetWriteDeadline() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
