package service

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGet_10(t *testing.T) {
	manager := &RoundTripperManager{
		roundTrippers: make(map[string]http.RoundTripper),
	}

	manager.roundTrippers["default@internal"] = &http.Transport{}

	tests := []struct {
		name     string
		input    string
		expected http.RoundTripper
		err      error
	}{
		{
			name:     "get existing roundtripper",
			input:    "default@internal",
			expected: &http.Transport{},
			err:      nil,
		},
		{
			name:     "get non-existing roundtripper",
			input:    "non-existing",
			expected: nil,
			err:      fmt.Errorf("servers transport not found non-existing"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rt, err := manager.Get(test.input)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("expected error %v, got %v", test.err, err)
			}

			if rt != test.expected {
				t.Errorf("expected roundtripper %v, got %v", test.expected, rt)
			}
		})
	}
}
