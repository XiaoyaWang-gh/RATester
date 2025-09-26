package connection

import (
	"fmt"
	"os"
	"testing"
)

func TestInitConnectionService_1(t *testing.T) {
	tests := []struct {
		name     string
		env      map[string]string
		expected string
	}{
		{
			name: "Test case 1",
			env: map[string]string{
				"bridge_port":      "8080",
				"https_proxy_port": "8081",
				"http_proxy_port":  "8082",
				"web_port":         "8083",
			},
			expected: "8080",
		},
		{
			name: "Test case 2",
			env: map[string]string{
				"bridge_port":      "8084",
				"https_proxy_port": "8085",
				"http_proxy_port":  "8086",
				"web_port":         "8087",
			},
			expected: "8084",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			for k, v := range tt.env {
				os.Setenv(k, v)
			}

			InitConnectionService()

			if bridgePort != tt.expected {
				t.Errorf("Expected bridgePort to be %s, got %s", tt.expected, bridgePort)
			}

			for k := range tt.env {
				os.Unsetenv(k)
			}
		})
	}
}
