package postpub

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestNewPostPublishResource_1(t *testing.T) {
	type args struct {
		id int
		r  resource.Resource
	}
	tests := []struct {
		name string
		args args
		want PostPublishedResource
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

			got := NewPostPublishResource(tt.args.id, tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostPublishResource() = %v, want %v", got, tt.want)
			}
		})
	}
}
