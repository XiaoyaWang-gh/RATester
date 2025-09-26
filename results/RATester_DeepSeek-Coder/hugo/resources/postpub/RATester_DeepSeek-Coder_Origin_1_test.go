package postpub

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestOrigin_1(t *testing.T) {
	type fields struct {
		prefix   string
		delegate resource.Resource
	}
	tests := []struct {
		name   string
		fields fields
		want   resource.Resource
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

			r := &PostPublishResource{
				prefix:   tt.fields.prefix,
				delegate: tt.fields.delegate,
			}
			if got := r.Origin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostPublishResource.Origin() = %v, want %v", got, tt.want)
			}
		})
	}
}
