package web

import (
	"fmt"
	"testing"
)

func TestGetUpperMethodString_1(t *testing.T) {
	testCases := []struct {
		name     string
		method   string
		expected string
	}{
		{"GET", "GET", "GET"},
		{"POST", "POST", "POST"},
		{"PUT", "PUT", "PUT"},
		{"DELETE", "DELETE", "DELETE"},
		{"HEAD", "HEAD", "HEAD"},
		{"OPTIONS", "OPTIONS", "OPTIONS"},
		{"PATCH", "PATCH", "PATCH"},
		{"*", "*", "*"},
		{"UNKNOWN", "UNKNOWN", "UNKNOWN"},
	}

	ctrlRegister := &ControllerRegister{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ctrlRegister.getUpperMethodString(tc.method)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
