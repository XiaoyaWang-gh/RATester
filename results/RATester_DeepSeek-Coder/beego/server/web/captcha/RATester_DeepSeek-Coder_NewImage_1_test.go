package captcha

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewImage_1(t *testing.T) {
	type args struct {
		digits []byte
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want *Image
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

			if got := NewImage(tt.args.digits, tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
