package images

import (
	"fmt"
	"image"
	"reflect"
	"testing"
)

func TestConfig_1(t *testing.T) {
	type args struct {
		path any
	}
	tests := []struct {
		name    string
		ns      *Namespace
		args    args
		want    image.Config
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

			got, err := tt.ns.Config(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Config() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Config() = %v, want %v", got, tt.want)
			}
		})
	}
}
