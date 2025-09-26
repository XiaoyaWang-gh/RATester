package client

import (
	"fmt"
	"testing"
)

func TestCookie_2(t *testing.T) {
	type test struct {
		name     string
		r        *Request
		key      string
		expected string
	}

	tests := []test{
		{
			name: "TestCookie_0",
			r: &Request{
				cookies: &Cookie{
					"key1": "value1",
					"key2": "value2",
				},
			},
			key:      "key1",
			expected: "value1",
		},
		{
			name: "TestCookie_1",
			r: &Request{
				cookies: &Cookie{
					"key1": "value1",
					"key2": "value2",
				},
			},
			key:      "key3",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.r.Cookie(tt.key); got != tt.expected {
				t.Errorf("Request.Cookie() = %v, want %v", got, tt.expected)
			}
		})
	}
}
