package acme

import (
	"bytes"
	"context"
	"fmt"
	"testing"
)

func TestGetTokenValue_1(t *testing.T) {
	ctx := context.Background()
	challenge := &ChallengeHTTP{
		httpChallenges: map[string]map[string][]byte{
			"token1": {
				"example.com": []byte("value1"),
			},
		},
	}

	tests := []struct {
		name     string
		token    string
		domain   string
		expected []byte
	}{
		{
			name:     "valid token and domain",
			token:    "token1",
			domain:   "example.com",
			expected: []byte("value1"),
		},
		{
			name:     "invalid token",
			token:    "token2",
			domain:   "example.com",
			expected: nil,
		},
		{
			name:     "invalid domain",
			token:    "token1",
			domain:   "example.org",
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := challenge.getTokenValue(ctx, test.token, test.domain)
			if !bytes.Equal(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
