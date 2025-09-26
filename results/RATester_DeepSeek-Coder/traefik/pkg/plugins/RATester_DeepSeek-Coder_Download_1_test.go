package plugins

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestDownload_1(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx      context.Context
		pName    string
		pVersion string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestDownload_Success",
			args: args{
				ctx:      context.Background(),
				pName:    "test",
				pVersion: "1.0.0",
			},
			want:    "hash",
			wantErr: false,
		},
		{
			name: "TestDownload_Error",
			args: args{
				ctx:      context.Background(),
				pName:    "error",
				pVersion: "1.0.0",
			},
			want:    "",
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
			got, err := c.Download(tt.args.ctx, tt.args.pName, tt.args.pVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Download() got = %v, want %v", got, tt.want)
			}
		})
	}
}
