package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestCloneWithMetadataFromMapIfNeeded_1(t *testing.T) {
	type args struct {
		m []map[string]any
		r resource.Resource
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

			if got := CloneWithMetadataFromMapIfNeeded(tt.args.m, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloneWithMetadataFromMapIfNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}
