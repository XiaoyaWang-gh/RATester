package filenotify

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

func TestAdd_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	poller := &filePoller{
		interval: 1 * time.Second,
		watches:  make(map[string]struct{}),
		done:     make(chan struct{}),
		events:   make(chan fsnotify.Event),
		errors:   make(chan error),
		mu:       sync.Mutex{},
		closed:   false,
	}

	testCases := []struct {
		name     string
		expected error
	}{
		{"test.txt", nil},
		{"", os.ErrNotExist},
		{"test.txt", fmt.Errorf("watch exists")},
	}

	for _, tc := range testCases {
		err := poller.Add(tc.name)
		if err != tc.expected {
			t.Errorf("Expected error %v, got %v", tc.expected, err)
		}
	}
}
