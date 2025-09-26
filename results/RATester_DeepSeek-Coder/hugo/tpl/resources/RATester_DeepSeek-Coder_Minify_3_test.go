package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestMinify_3(t *testing.T) {
	type args struct {
		r resources.ResourceTransformer
	}
	tests := []struct {
		name    string
		ns      *Namespace
		args    args
		want    resource.Resource
		wantErr bool
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

			got, err := tt.ns.Minify(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Minify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Minify() = %v, want %v", got, tt.want)
			}
		})
	}
}
