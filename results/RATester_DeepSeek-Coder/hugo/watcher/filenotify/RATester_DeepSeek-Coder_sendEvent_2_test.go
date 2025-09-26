package filenotify

import (
	"fmt"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

func TestSendEvent_2(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		event   fsnotify.Event
		wantErr bool
	}{
		{
			name: "send event",
			event: fsnotify.Event{
				Name: "test.txt",
				Op:   fsnotify.Create,
			},
			wantErr: false,
		},
		{
			name: "send event with error",
			event: fsnotify.Event{
				Name: "test.txt",
				Op:   fsnotify.Create,
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fp := &filePoller{
				events: make(chan fsnotify.Event),
				done:   make(chan struct{}),
			}

			go func() {
				select {
				case e := <-fp.Events():
					if e.Name != tc.event.Name || e.Op != tc.event.Op {
						t.Errorf("Expected event %v, got %v", tc.event, e)
					}
				case <-time.After(1 * time.Second):
					t.Errorf("Timeout waiting for event")
				}
			}()

			err := fp.sendEvent(tc.event)
			if (err != nil) != tc.wantErr {
				t.Errorf("Expected error %v, got %v", tc.wantErr, err != nil)
			}

			close(fp.done)
		})
	}
}
