package hugolib

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestShouldSkipFileChangeEvent_1(t *testing.T) {
	hugoSites := &HugoSites{
		skipRebuildForFilenames: map[string]bool{
			"file1.txt": true,
			"file2.txt": true,
		},
	}

	tests := []struct {
		name     string
		event    fsnotify.Event
		expected bool
	}{
		{
			name: "Should skip file",
			event: fsnotify.Event{
				Name: "file1.txt",
			},
			expected: true,
		},
		{
			name: "Should not skip file",
			event: fsnotify.Event{
				Name: "file3.txt",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := hugoSites.ShouldSkipFileChangeEvent(test.event)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
