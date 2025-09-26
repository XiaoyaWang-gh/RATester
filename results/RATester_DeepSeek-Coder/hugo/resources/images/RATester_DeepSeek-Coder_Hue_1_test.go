package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestHue_1(t *testing.T) {
	type args struct {
		shift any
	}
	tests := []struct {
		name string
		args args
		want gift.Filter
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

			f := &Filters{}
			if got := f.Hue(tt.args.shift); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue() = %v, want %v", got, tt.want)
			}
		})
	}
}
