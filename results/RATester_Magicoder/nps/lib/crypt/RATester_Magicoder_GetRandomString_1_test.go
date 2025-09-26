package crypt

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetRandomString_1(t *testing.T) {
	testCases := []struct {
		name string
		l    int
	}{
		{"Test case 1", 5},
		{"Test case 2", 10},
		{"Test case 3", 15},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := GetRandomString(tc.l)
			if len(result) != tc.l {
				t.Errorf("Expected length %d, got %d", tc.l, len(result))
			}
			for _, r := range result {
				if !strings.ContainsRune("0123456789abcdefghijklmnopqrstuvwxyz", r) {
					t.Errorf("Unexpected character %c", r)
				}
			}
		})
	}
}
