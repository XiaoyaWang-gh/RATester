package filenotify

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Testwatch_1(t *testing.T) {
	type fields struct {
		interval time.Duration
		watches  map[string]struct{}
		done     chan struct{}
		events   chan fsnotify.Event
		errors   chan error
		mu       sync.Mutex
		closed   bool
	}
	type args struct {
		item *itemToWatch
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   func(t *testing.T, w *filePoller)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := &filePoller{
				interval: tt.fields.interval,
				watches:  tt.fields.watches,
				done:     tt.fields.done,
				events:   tt.fields.events,
				errors:   tt.fields.errors,
				mu:       tt.fields.mu,
				closed:   tt.fields.closed,
			}
			w.watch(tt.args.item)
			tt.want(t, w)
		})
	}
}
