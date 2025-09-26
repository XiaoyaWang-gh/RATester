package js

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/resources"
)

func TestNew_26(t *testing.T) {
	type args struct {
		fs *filesystems.SourceFilesystem
		rs *resources.Spec
	}
	tests := []struct {
		name string
		args args
		want *Client
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

			if got := New(tt.args.fs, tt.args.rs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
