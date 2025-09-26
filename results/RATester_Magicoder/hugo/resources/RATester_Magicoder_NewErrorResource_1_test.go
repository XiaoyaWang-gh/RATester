package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestNewErrorResource_1(t *testing.T) {
	type args struct {
		err resource.ResourceError
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

			if got := NewErrorResource(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorResource() = %v, want %v", got, tt.want)
			}
		})
	}
}
