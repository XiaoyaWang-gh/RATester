package static

import (
	"fmt"
	"testing"
)

func TestGetAddress_1(t *testing.T) {
	tests := []struct {
		name     string
		ep       EntryPoint
		expected string
	}{
		{
			name: "should return the first part of the address",
			ep: EntryPoint{
				Address: "127.0.0.1:8080",
			},
			expected: "127.0.0.1",
		},
		{
			name: "should return the first part of the address even if there are more slashes",
			ep: EntryPoint{
				Address: "127.0.0.1/tcp",
			},
			expected: "127.0.0.1",
		},
		{
			name: "should return the first part of the address even if there are more slashes and a port",
			ep: EntryPoint{
				Address: "127.0.0.1:8080/tcp",
			},
			expected: "127.0.0.1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.ep.GetAddress()
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
