package cssjs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestProcess_3(t *testing.T) {
	type args struct {
		res     resources.ResourceTransformer
		options map[string]any
	}
	tests := []struct {
		name    string
		c       *TailwindCSSClient
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

			got, err := tt.c.Process(tt.args.res, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("TailwindCSSClient.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TailwindCSSClient.Process() = %v, want %v", got, tt.want)
			}
		})
	}
}
