package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestFromString_3(t *testing.T) {
	type args struct {
		targetPathIn any
		contentIn    any
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

			ns := &Namespace{}
			got, err := ns.FromString(tt.args.targetPathIn, tt.args.contentIn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
