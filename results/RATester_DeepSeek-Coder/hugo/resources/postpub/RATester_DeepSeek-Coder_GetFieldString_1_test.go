package postpub

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestGetFieldString_1(t *testing.T) {
	type fields struct {
		prefix   string
		delegate resource.Resource
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
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
			got, got1 := r.GetFieldString(tt.args.pattern)
			if got != tt.want {
				t.Errorf("GetFieldString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetFieldString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
