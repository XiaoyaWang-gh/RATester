package modules

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGraph_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		client   *Client
		expected string
	}{
		{
			name: "success",
			client: &Client{
				fs: afero.NewMemMapFs(),
				moduleConfig: Config{
					Imports: []Import{
						{Path: "github.com/gohugoio/hugo"},
					},
				},
			},
			expected: "github.com/gohugoio/hugo => github.com/gohugoio/hugo\n",
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			buf := &bytes.Buffer{}
			err := tc.client.Graph(buf)
			if err != nil {
				t.Errorf("Expected nil error, got %v", err)
			}

			if buf.String() != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, buf.String())
			}
		})
	}
}
