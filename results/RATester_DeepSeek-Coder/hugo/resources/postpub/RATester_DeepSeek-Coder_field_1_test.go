package postpub

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestField_1(t *testing.T) {
	type fields struct {
		prefix   string
		delegate resource.Resource
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
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
			if got := r.field(tt.args.name); got != tt.want {
				t.Errorf("PostPublishResource.field() = %v, want %v", got, tt.want)
			}
		})
	}
}
