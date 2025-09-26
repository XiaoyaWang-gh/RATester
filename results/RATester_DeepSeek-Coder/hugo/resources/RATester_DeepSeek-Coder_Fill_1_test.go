package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/images"
)

func TestFill_1(t *testing.T) {
	type args struct {
		spec string
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

			got, err := tt.i.Fill(tt.args.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("imageResource.Fill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageResource.Fill() = %v, want %v", got, tt.want)
			}
		})
	}
}
