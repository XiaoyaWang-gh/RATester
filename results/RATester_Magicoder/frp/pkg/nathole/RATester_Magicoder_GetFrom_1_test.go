package nathole

import (
	"fmt"
	"testing"

	"github.com/pion/stun/v2"
)

func TestGetFrom_1(t *testing.T) {
	testCases := []struct {
		name    string
		message *stun.Message
		wantErr bool
	}{
		{
			name:    "success",
			message: &stun.Message{},
			wantErr: false,
		},
		{
			name:    "failure",
			message: &stun.Message{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			changedAddress := &ChangedAddress{}
			err := changedAddress.GetFrom(tc.message)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetFrom() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
