package gin

import (
	"fmt"
	"testing"
)

func TestvalidateHeader_1(t *testing.T) {
	engine := &Engine{
		trustedProxies: []string{"192.168.0.1"},
	}

	tests := []struct {
		name     string
		header   string
		expected string
		valid    bool
	}{
		{
			name:     "valid header",
			header:   "192.168.0.1, 192.168.0.2",
			expected: "192.168.0.1",
			valid:    true,
		},
		{
			name:     "invalid header",
			header:   "192.168.0.3, 192.168.0.4",
			expected: "",
			valid:    false,
		},
		{
			name:     "empty header",
			header:   "",
			expected: "",
			valid:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			clientIP, valid := engine.validateHeader(test.header)
			if clientIP != test.expected || valid != test.valid {
				t.Errorf("validateHeader(%s) = %s, %t; want %s, %t", test.header, clientIP, valid, test.expected, test.valid)
			}
		})
	}
}
