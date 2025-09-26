package httplib

import (
	"context"
	"fmt"
	"testing"
)

func TestDoRequestWithCtx_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		request *BeegoHTTPRequest
		ctx     context.Context
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "success",
			request: &BeegoHTTPRequest{},
			ctx:     context.Background(),
			wantErr: false,
		},
		{
			name:    "error",
			request: &BeegoHTTPRequest{},
			ctx:     context.Background(),
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

			_, err := tc.request.DoRequestWithCtx(tc.ctx)
			if (err != nil) != tc.wantErr {
				t.Errorf("DoRequestWithCtx() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
