package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/internal"
)

func TestGetResourcePaths_1(t *testing.T) {
	type fields struct {
		paths internal.ResourcePaths
	}
	tests := []struct {
		name   string
		fields fields
		want   internal.ResourcePaths
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

			l := &genericResource{
				paths: tt.fields.paths,
			}
			if got := l.getResourcePaths(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genericResource.getResourcePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
