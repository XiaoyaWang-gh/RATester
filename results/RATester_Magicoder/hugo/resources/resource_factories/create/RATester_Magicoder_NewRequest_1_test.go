package create

import (
	"fmt"
	"testing"
)

func TestNewRequest_1(t *testing.T) {
	testCases := []struct {
		name    string
		options fromRemoteOptions
		url     string
		wantErr bool
	}{
		{
			name: "success",
			options: fromRemoteOptions{
				Method: "GET",
				Headers: map[string]any{
					"Content-Type": "application/json",
				},
				Body: []byte(`{"key": "value"}`),
			},
			url:     "http://example.com",
			wantErr: false,
		},
		{
			name: "failure",
			options: fromRemoteOptions{
				Method: "GET",
				Headers: map[string]any{
					"Content-Type": "application/json",
				},
				Body: []byte(`{"key": "value"}`),
			},
			url:     ":",
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

			_, err := tc.options.NewRequest(tc.url)
			if (err != nil) != tc.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
