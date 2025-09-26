package server

import (
	"errors"
	"fmt"
	"testing"
)

func TestStopServer_1(t *testing.T) {
	testCases := []struct {
		name     string
		id       int
		expected error
	}{
		{"stop running server", 1, nil},
		{"stop not running server", 2, errors.New("task is not running")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual := StopServer(tc.id)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
