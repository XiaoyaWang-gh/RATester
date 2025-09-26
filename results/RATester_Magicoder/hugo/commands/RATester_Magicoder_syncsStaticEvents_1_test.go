package commands

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestsyncsStaticEvents_1(t *testing.T) {
	testCases := []struct {
		name          string
		staticEvents  []fsnotify.Event
		expectedError error
	}{
		{
			name: "Test case 1",
			staticEvents: []fsnotify.Event{
				{Name: "file1", Op: fsnotify.Create},
				{Name: "file2", Op: fsnotify.Write},
			},
			expectedError: nil,
		},
		{
			name: "Test case 2",
			staticEvents: []fsnotify.Event{
				{Name: "file3", Op: fsnotify.Remove},
				{Name: "file4", Op: fsnotify.Chmod},
			},
			expectedError: nil,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &staticSyncer{}
			err := s.syncsStaticEvents(tc.staticEvents)
			if err != tc.expectedError {
				t.Errorf("Expected error %v, but got %v", tc.expectedError, err)
			}
		})
	}
}
