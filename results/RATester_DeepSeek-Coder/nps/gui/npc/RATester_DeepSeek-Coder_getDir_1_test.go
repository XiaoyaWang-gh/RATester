package main

import (
	"fmt"
	"testing"
)

func TestGetDir_1(t *testing.T) {
	type test struct {
		name     string
		expected string
		wantErr  bool
	}

	tests := []test{
		{
			name:     "Test Case 1",
			expected: "/data/data/org.nps.client/files",
			wantErr:  false,
		},
		{
			name:     "Test Case 2",
			expected: "/home/user/.config",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := getDir()
			if (err != nil) != tt.wantErr {
				t.Errorf("getDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("getDir() = %v, want %v", got, tt.expected)
			}
		})
	}
}
