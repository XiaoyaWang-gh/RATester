package gin

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestParseIP_1(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected net.IP
	}

	tests := []test{
		{
			name:     "IPv4",
			input:    "192.168.1.1",
			expected: net.IPv4(192, 168, 1, 1),
		},
		{
			name:     "IPv6",
			input:    "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expected: net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
		},
		{
			name:     "Invalid IP",
			input:    "invalid",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := parseIP(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("parseIP(%s) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
