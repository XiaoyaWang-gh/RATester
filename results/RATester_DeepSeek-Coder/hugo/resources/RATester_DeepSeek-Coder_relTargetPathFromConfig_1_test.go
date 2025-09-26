package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/images"
	"github.com/gohugoio/hugo/resources/internal"
)

func TestRelTargetPathFromConfig_1(t *testing.T) {
	tests := []struct {
		name string
		i    *imageResource
		conf images.ImageConfig
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

			if got := tt.i.relTargetPathFromConfig(tt.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageResource.relTargetPathFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
