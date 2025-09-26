package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/cache/dynacache"
)

func TestnewResourceCache_1(t *testing.T) {
	type args struct {
		rs       *Spec
		memCache *dynacache.Cache
	}
	tests := []struct {
		name string
		args args
		want *ResourceCache
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

			if got := newResourceCache(tt.args.rs, tt.args.memCache); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newResourceCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
