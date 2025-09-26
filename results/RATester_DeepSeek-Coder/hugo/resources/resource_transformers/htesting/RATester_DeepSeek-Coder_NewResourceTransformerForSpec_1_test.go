package htesting

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestNewResourceTransformerForSpec_1(t *testing.T) {
	type args struct {
		spec     *resources.Spec
		filename string
		content  string
	}
	tests := []struct {
		name    string
		args    args
		want    resources.ResourceTransformer
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

			got, err := NewResourceTransformerForSpec(tt.args.spec, tt.args.filename, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResourceTransformerForSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourceTransformerForSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
