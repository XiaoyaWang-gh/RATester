package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/internal"
)

func TestsetTargetPath_1(t *testing.T) {
	type args struct {
		d internal.ResourcePaths
	}
	tests := []struct {
		name string
		args args
		want internal.ResourcePaths
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

			l := &genericResource{}
			l.setTargetPath(tt.args.d)
			if !reflect.DeepEqual(l.paths, tt.want) {
				t.Errorf("setTargetPath() = %v, want %v", l.paths, tt.want)
			}
		})
	}
}
