package commands

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/gohugoio/hugo/watcher"
)

func TestHandleEvents_2(t *testing.T) {
	type args struct {
		watcher      *watcher.Batcher
		staticSyncer *staticSyncer
		evs          []fsnotify.Event
		configSet    map[string]bool
	}
	tests := []struct {
		name string
		args args
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

			c := &hugoBuilder{}
			c.handleEvents(tt.args.watcher, tt.args.staticSyncer, tt.args.evs, tt.args.configSet)
		})
	}
}
