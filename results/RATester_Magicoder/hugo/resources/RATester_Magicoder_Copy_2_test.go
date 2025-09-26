package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestCopy_2(t *testing.T) {
	type args struct {
		r          resource.Resource
		targetPath string
	}
	tests := []struct {
		name string
		args args
		want resource.Resource
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

			if got := Copy(tt.args.r, tt.args.targetPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
