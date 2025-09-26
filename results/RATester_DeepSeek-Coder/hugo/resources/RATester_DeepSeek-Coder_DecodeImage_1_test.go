package resources

import (
	"fmt"
	"image"
	"reflect"
	"testing"
)

func TestDecodeImage_1(t *testing.T) {
	tests := []struct {
		name    string
		i       *imageResource
		want    image.Image
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

			got, err := tt.i.DecodeImage()
			if (err != nil) != tt.wantErr {
				t.Errorf("imageResource.DecodeImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageResource.DecodeImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
