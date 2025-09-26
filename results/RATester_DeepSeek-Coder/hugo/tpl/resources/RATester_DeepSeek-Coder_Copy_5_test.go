package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestCopy_5(t *testing.T) {
	type args struct {
		s any
		r resource.Resource
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

			got, err := tt.ns.Copy(tt.args.s, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
