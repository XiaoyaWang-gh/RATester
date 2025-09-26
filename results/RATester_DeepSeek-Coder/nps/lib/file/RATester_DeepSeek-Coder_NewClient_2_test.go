package file

import (
	"fmt"
	"testing"
)

func TestNewClient_2(t *testing.T) {
	db := &DbUtils{
		JsonDb: &JsonDb{},
	}

	testCases := []struct {
		name    string
		input   *Client
		wantErr bool
	}{
		{
			name: "success",
			input: &Client{
				WebUserName: "test",
				VerifyKey:   "test",
			},
			wantErr: false,
		},
		{
			name: "duplicate username",
			input: &Client{
				WebUserName: "test",
				VerifyKey:   "test",
			},
			wantErr: true,
		},
		{
			name: "duplicate verify key",
			input: &Client{
				WebUserName: "test2",
				VerifyKey:   "test",
			},
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

			err := db.NewClient(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
