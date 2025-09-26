package images

import (
	"fmt"
	"image"
	"reflect"
	"testing"
)

func TestApplyFiltersFromConfig_1(t *testing.T) {
	type args struct {
		src  image.Image
		conf ImageConfig
	}
	tests := []struct {
		name    string
		args    args
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

			p := &ImageProcessor{}
			got, err := p.ApplyFiltersFromConfig(tt.args.src, tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyFiltersFromConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyFiltersFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
