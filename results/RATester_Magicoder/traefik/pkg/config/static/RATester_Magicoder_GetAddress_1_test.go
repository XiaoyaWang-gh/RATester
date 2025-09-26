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
				Address: "192.168.1.1:80",
			},
			expected: "192.168.1.1",
		},
		{
			name: "should return the first part of the address even if there is a path",
			ep: EntryPoint{
				Address: "192.168.1.1:80/path",
			},
			expected: "192.168.1.1",
		},
		{
			name: "should return the first part of the address even if there is a path and a query",
			ep: EntryPoint{
				Address: "192.168.1.1:80/path?query=test",
			},
			expected: "192.168.1.1",
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
