package resources

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestTryTransformedFileCache_1(t *testing.T) {
	type args struct {
		key string
		u   *transformationUpdate
	}
	tests := []struct {
		name string
		r    *genericResource
		args args
		want io.ReadCloser
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

			if got := tt.r.tryTransformedFileCache(tt.args.key, tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genericResource.tryTransformedFileCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
