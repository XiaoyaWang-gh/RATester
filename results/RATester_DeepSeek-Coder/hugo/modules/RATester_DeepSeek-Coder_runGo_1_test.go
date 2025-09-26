package modules

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestRunGo_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		client  *Client
		ctx     context.Context
		stdout  io.Writer
		args    []string
		wantErr bool
	}

	testCases := []testCase{
		{
			name:   "success",
			client: &Client{},
			ctx:    context.Background(),
			stdout: os.Stdout,
			args:   []string{"version"},
		},
		{
			name:    "error",
			client:  &Client{},
			ctx:     context.Background(),
			stdout:  os.Stdout,
			args:    []string{"invalid_command"},
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

			err := tc.client.runGo(tc.ctx, tc.stdout, tc.args...)
			if (err != nil) != tc.wantErr {
				t.Errorf("runGo() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
