package plugins

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestCheck_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		ctx      context.Context
		pName    string
		pVersion string
		hash     string
		wantErr  bool
	}{
		{
			name:     "success",
			ctx:      context.Background(),
			pName:    "test",
			pVersion: "1.0.0",
			hash:     "testHash",
			wantErr:  false,
		},
		{
			name:     "failure",
			ctx:      context.Background(),
			pName:    "test",
			pVersion: "1.0.0",
			hash:     "wrongHash",
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

			c := &Client{
				HTTPClient: &http.Client{
					Timeout: 10 * time.Second,
				},
				baseURL: &url.URL{
					Scheme: "http",
					Host:   "localhost:8080",
				},
			}

			err := c.Check(tt.ctx, tt.pName, tt.pVersion, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
