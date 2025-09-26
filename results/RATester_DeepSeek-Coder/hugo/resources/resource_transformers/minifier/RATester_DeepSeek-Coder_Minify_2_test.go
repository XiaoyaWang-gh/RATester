package minifier

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestMinify_2(t *testing.T) {
	type args struct {
		res resources.ResourceTransformer
	}
	tests := []struct {
		name    string
		c       *Client
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

			got, err := tt.c.Minify(tt.args.res)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Minify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Minify() = %v, want %v", got, tt.want)
			}
		})
	}
}
