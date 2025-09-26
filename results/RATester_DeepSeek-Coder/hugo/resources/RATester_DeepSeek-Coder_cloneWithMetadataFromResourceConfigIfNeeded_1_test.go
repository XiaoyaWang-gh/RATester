package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page/pagemeta"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestCloneWithMetadataFromResourceConfigIfNeeded_1(t *testing.T) {
	type args struct {
		rc *pagemeta.ResourceConfig
		r  resource.Resource
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

			if got := cloneWithMetadataFromResourceConfigIfNeeded(tt.args.rc, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneWithMetadataFromResourceConfigIfNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}
