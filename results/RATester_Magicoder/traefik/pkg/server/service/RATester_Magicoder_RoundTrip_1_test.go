package service

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRoundTrip_1(t *testing.T) {
	tests := []struct {
		name    string
		req     *http.Request
		wantErr bool
	}{
		{
			name: "Test with Connection Upgrade",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"Upgrade"},
				},
			},
			wantErr: false,
		},
		{
			name: "Test without Connection Upgrade",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"Close"},
				},
			},
			wantErr: false,
		},
		{
			name:    "Test with nil Request",
			req:     nil,
			wantErr: true,
		},
	}

	rt := &smartRoundTripper{
		http:  &http.Transport{},
		http2: &http.Transport{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := rt.RoundTrip(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoundTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
