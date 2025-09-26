package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/images"
)

func TestProcessActionSpec_1(t *testing.T) {
	type args struct {
		action string
		spec   string
	}
	tests := []struct {
		name    string
		i       *imageResource
		args    args
		want    images.ImageResource
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

			got, err := tt.i.processActionSpec(tt.args.action, tt.args.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("imageResource.processActionSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageResource.processActionSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
