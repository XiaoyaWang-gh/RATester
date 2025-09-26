package session

import (
	"fmt"
	"testing"
)

func TestGenerateRandomKey_1(t *testing.T) {
	tests := []struct {
		name     string
		strength int
		want     int
	}{
		{"Test Case 1", 16, 16},
		{"Test Case 2", 32, 32},
		{"Test Case 3", 64, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := generateRandomKey(tt.strength)
			if len(got) != tt.want {
				t.Errorf("generateRandomKey() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
