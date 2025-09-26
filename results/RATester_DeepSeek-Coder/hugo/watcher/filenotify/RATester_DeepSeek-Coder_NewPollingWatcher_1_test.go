package filenotify

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

func TestNewPollingWatcher_1(t *testing.T) {
	type args struct {
		interval time.Duration
	}
	tests := []struct {
		name string
		args args
		want FileWatcher
	}{
		{
			name: "Test case 1",
			args: args{
				interval: 1 * time.Second,
			},
			want: &filePoller{
				interval: 1 * time.Second,
				done:     make(chan struct{}),
				events:   make(chan fsnotify.Event),
				errors:   make(chan error),
			},
		},
		{
			name: "Test case 2",
			args: args{
				interval: 2 * time.Second,
			},
			want: &filePoller{
				interval: 2 * time.Second,
				done:     make(chan struct{}),
				events:   make(chan fsnotify.Event),
				errors:   make(chan error),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewPollingWatcher(tt.args.interval)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPollingWatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
