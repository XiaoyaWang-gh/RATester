package livereloadinject

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/transform"
)

func TestNew_14(t *testing.T) {
	type args struct {
		baseURL *url.URL
	}
	tests := []struct {
		name string
		args args
		want transform.Transformer
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

			if got := New(tt.args.baseURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
