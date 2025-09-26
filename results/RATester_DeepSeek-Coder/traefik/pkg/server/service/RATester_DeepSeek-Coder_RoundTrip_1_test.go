package service

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRoundTrip_1(t *testing.T) {
	type testCase struct {
		name       string
		req        *http.Request
		wantResp   *http.Response
		wantErr    bool
		wantCloser bool
	}

	tests := []testCase{
		{
			name: "Upgrade connection",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"Upgrade"},
				},
			},
			wantErr:    false,
			wantCloser: true,
		},
		{
			name: "No upgrade connection",
			req: &http.Request{
				Header: http.Header{
					"Connection": []string{"Close"},
				},
			},
			wantErr:    false,
			wantCloser: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &smartRoundTripper{
				http:  &http.Transport{},
				http2: &http.Transport{},
			}
			gotResp, err := m.RoundTrip(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("smartRoundTripper.RoundTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotResp != nil) != tt.wantCloser {
				t.Errorf("smartRoundTripper.RoundTrip() gotCloser = %v, wantCloser %v", gotResp != nil, tt.wantCloser)
			}
		})
	}
}
