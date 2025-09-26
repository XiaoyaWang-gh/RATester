package images

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TesttoColorGo_1(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    color.Color
		want1   bool
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

			got, got1, err := toColorGo(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("toColorGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toColorGo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("toColorGo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
