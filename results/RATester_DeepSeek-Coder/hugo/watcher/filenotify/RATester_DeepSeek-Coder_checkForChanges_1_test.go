package filenotify

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestCheckForChanges_1(t *testing.T) {
	type args struct {
		item *itemToWatch
	}
	tests := []struct {
		name    string
		args    args
		want    []fsnotify.Event
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

			got, err := tt.args.item.checkForChanges()
			if (err != nil) != tt.wantErr {
				t.Errorf("checkForChanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkForChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
