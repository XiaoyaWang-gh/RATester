package proxy

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConfigDefault_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		config   Config
		expected Config
	}{
		{
			name: "Test with empty config",
			config: Config{
				Servers: []string{"https://foobar.com", "http://www.foobar.com"},
			},
			expected: Config{
				Servers: []string{"https://foobar.com", "http://www.foobar.com"},
				Timeout: ConfigDefault.Timeout,
			},
		},
		{
			name: "Test with custom timeout",
			config: Config{
				Servers: []string{"https://foobar.com", "http://www.foobar.com"},
				Timeout: 5 * time.Second,
			},
			expected: Config{
				Servers: []string{"https://foobar.com", "http://www.foobar.com"},
				Timeout: 5 * time.Second,
			},
		},
		{
			name: "Test with no servers",
			config: Config{
				Timeout: 5 * time.Second,
			},
			expected: Config{
				Timeout: 5 * time.Second,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := configDefault(test.config)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
