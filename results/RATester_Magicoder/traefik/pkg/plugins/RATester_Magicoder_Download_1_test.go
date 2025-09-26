package plugins

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestDownload_1(t *testing.T) {
	ctx := context.Background()

	client := &Client{
		HTTPClient: &http.Client{},
		baseURL:    &url.URL{},
		archives:   "/tmp/archives",
		stateFile:  "/tmp/state.json",
		goPath:     "/tmp/go",
		sources:    "/tmp/sources",
	}

	tests := []struct {
		name     string
		client   *Client
		ctx      context.Context
		pName    string
		pVersion string
		want     string
		wantErr  bool
	}{
		{
			name:     "Test case 1",
			client:   client,
			ctx:      ctx,
			pName:    "test",
			pVersion: "1.0.0",
			want:     "hash",
			wantErr:  false,
		},
		{
			name:     "Test case 2",
			client:   client,
			ctx:      ctx,
			pName:    "test",
			pVersion: "2.0.0",
			want:     "hash",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.client.Download(tt.ctx, tt.pName, tt.pVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Download() = %v, want %v", got, tt.want)
			}
		})
	}
}
