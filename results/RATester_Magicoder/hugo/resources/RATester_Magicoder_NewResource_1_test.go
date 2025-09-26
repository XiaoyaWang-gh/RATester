package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestNewResource_1(t *testing.T) {
	type args struct {
		rd ResourceSourceDescriptor
	}
	tests := []struct {
		name    string
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

			r := &Spec{}
			got, err := r.NewResource(tt.args.rd)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResource() = %v, want %v", got, tt.want)
			}
		})
	}
}
