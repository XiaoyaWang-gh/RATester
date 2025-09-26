package watcher

import (
	"fmt"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gohugoio/hugo/watcher/filenotify"
)

func TestRun_3(t *testing.T) {
	type fields struct {
		FileWatcher filenotify.FileWatcher
		ticker      *time.Ticker
		done        chan struct{}
		Events      chan []fsnotify.Event
	}
	tests := []struct {
		name   string
		fields fields
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

			b := &Batcher{
				FileWatcher: tt.fields.FileWatcher,
				ticker:      tt.fields.ticker,
				done:        tt.fields.done,
				Events:      tt.fields.Events,
			}
			b.run()
		})
	}
}
