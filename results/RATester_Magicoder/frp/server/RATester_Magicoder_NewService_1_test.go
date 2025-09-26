package server

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewService_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := &v1.ServerConfig{
		BindAddr: "127.0.0.1",
		BindPort: 7000,
		// other fields
	}
	svr, err := NewService(cfg)
	if err != nil {
		t.Fatalf("Failed to create service: %v", err)
	}

	// Test cases
	testCases := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{
			name:     "BindAddr",
			expected: "127.0.0.1",
			actual:   svr.cfg.BindAddr,
		},
		{
			name:     "BindPort",
			expected: 7000,
			actual:   svr.cfg.BindPort,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		if tc.expected != tc.actual {
			t.Errorf("Expected %v for %s, but got %v", tc.expected, tc.name, tc.actual)
		}
	}
}
