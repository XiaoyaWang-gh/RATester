package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/gohugoio/hugo/hugolib/filesystems"
)

func TestPartitionDynamicEvents_1(t *testing.T) {
	type args struct {
		sourceFs *filesystems.SourceFilesystems
		events   []fsnotify.Event
	}
	tests := []struct {
		name string
		args args
		want dynamicEvents
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

			if got := partitionDynamicEvents(tt.args.sourceFs, tt.args.events); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionDynamicEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
