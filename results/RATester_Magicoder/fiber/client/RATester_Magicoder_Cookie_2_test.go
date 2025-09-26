package client

import (
	"fmt"
	"testing"
)

func TestCookie_2(t *testing.T) {
	req := &Request{
		cookies: &Cookie{
			"key1": "value1",
			"key2": "value2",
		},
	}

	tests := []struct {
		name string
		key  string
		want string
	}{
		{"Test case 1", "key1", "value1"},
		{"Test case 2", "key2", "value2"},
		{"Test case 3", "key3", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := req.Cookie(tt.key); got != tt.want {
				t.Errorf("Cookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
