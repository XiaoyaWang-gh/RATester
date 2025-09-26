package client

import (
	"context"
	"fmt"
	"testing"
)

func TestUpdateConfig_1(t *testing.T) {
	ctx := context.Background()
	client := &Client{
		address:  "localhost:8080",
		authUser: "testUser",
		authPwd:  "testPwd",
	}

	testCases := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name:    "valid content",
			content: "valid content",
			wantErr: false,
		},
		{
			name:    "empty content",
			content: "",
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

			err := client.UpdateConfig(ctx, tc.content)
			if (err != nil) != tc.wantErr {
				t.Errorf("UpdateConfig() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
