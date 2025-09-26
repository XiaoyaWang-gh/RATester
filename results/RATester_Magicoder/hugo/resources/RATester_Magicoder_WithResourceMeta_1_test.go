package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestWithResourceMeta_1(t *testing.T) {
	type args struct {
		mp resource.ResourceMetaProvider
	}
	tests := []struct {
		name string
		r    resourceAdapter
		args args
		want *resourceAdapter
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

			if got := tt.r.WithResourceMeta(tt.args.mp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithResourceMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}
