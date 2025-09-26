package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/watcher"
)

func TestNewWatcher_1(t *testing.T) {
	type args struct {
		pollIntervalStr string
		dirList         []string
	}
	tests := []struct {
		name    string
		args    args
		want    *watcher.Batcher
		wantErr bool
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
			got, err := c.newWatcher(tt.args.pollIntervalStr, tt.args.dirList...)
			if (err != nil) != tt.wantErr {
				t.Errorf("hugoBuilder.newWatcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hugoBuilder.newWatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
