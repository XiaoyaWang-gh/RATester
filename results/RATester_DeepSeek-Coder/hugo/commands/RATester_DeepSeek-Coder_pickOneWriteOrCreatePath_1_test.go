package commands

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/gohugoio/hugo/config"
)

func TestPickOneWriteOrCreatePath_1(t *testing.T) {
	type args struct {
		contentTypes config.ContentTypesProvider
		events       []fsnotify.Event
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := pickOneWriteOrCreatePath(tt.args.contentTypes, tt.args.events); got != tt.want {
				t.Errorf("pickOneWriteOrCreatePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
