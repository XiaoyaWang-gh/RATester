package static

import (
	"fmt"
	"testing"
)

func TestGetProtocol_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    EntryPoint
		expected string
		err      error
	}

	tests := []test{
		{
			input: EntryPoint{
				Address: "tcp://localhost:80",
			},
			expected: "tcp",
			err:      nil,
		},
		{
			input: EntryPoint{
				Address: "udp://localhost:80",
			},
			expected: "udp",
			err:      nil,
		},
		{
			input: EntryPoint{
				Address: "http://localhost:80",
			},
			expected: "",
			err:      fmt.Errorf("invalid protocol: http"),
		},
		{
			input: EntryPoint{
				Address: "localhost:80",
			},
			expected: "tcp",
			err:      nil,
		},
	}

	for _, test := range tests {
		protocol, err := test.input.GetProtocol()
		if protocol != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, protocol)
		}
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error %s, got %s", test.err, err)
		}
	}
}
