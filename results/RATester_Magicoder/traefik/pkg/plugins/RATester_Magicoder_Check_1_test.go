package plugins

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestCheck_1(t *testing.T) {
	type args struct {
		ctx      context.Context
		pName    string
		pVersion string
		hash     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				ctx:      context.Background(),
				pName:    "test",
				pVersion: "1.0.0",
				hash:     "testHash",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				ctx:      context.Background(),
				pName:    "test",
				pVersion: "1.0.0",
				hash:     "",
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				ctx:      context.Background(),
				pName:    "test",
				pVersion: "1.0.0",
				hash:     "invalidHash",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Client{
				HTTPClient: &http.Client{},
				baseURL:    &url.URL{},
			}
			if err := c.Check(tt.args.ctx, tt.args.pName, tt.args.pVersion, tt.args.hash); (err != nil) != tt.wantErr {
				t.Errorf("Client.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
