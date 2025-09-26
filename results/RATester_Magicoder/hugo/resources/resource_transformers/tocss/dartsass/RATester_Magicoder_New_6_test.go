package dartsass

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/resources"
)

func TestNew_6(t *testing.T) {
	type args struct {
		fs *filesystems.SourceFilesystem
		rs *resources.Spec
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
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

			got, err := New(tt.args.fs, tt.args.rs)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
