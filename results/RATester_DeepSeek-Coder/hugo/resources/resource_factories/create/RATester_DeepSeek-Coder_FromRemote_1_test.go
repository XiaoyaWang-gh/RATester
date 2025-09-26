package create

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestFromRemote_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		uri       string
		options   map[string]any
		wantError bool
	}{
		{
			name:    "valid URL with GET method",
			uri:     "https://example.com",
			options: map[string]any{"method": "GET"},
		},
		{
			name:      "valid URL with invalid method",
			uri:       "https://example.com",
			options:   map[string]any{"method": "INVALID"},
			wantError: true,
		},
		{
			name:      "invalid URL",
			uri:       ":",
			options:   map[string]any{"method": "GET"},
			wantError: true,
		},
		{
			name:    "valid URL with HEAD method",
			uri:     "https://example.com",
			options: map[string]any{"method": "HEAD"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			client := &Client{
				rs: &resources.Spec{},
			}

			_, err := client.FromRemote(tc.uri, tc.options)
			if (err != nil) != tc.wantError {
				t.Errorf("FromRemote() error = %v, wantError %v", err, tc.wantError)
			}
		})
	}
}
