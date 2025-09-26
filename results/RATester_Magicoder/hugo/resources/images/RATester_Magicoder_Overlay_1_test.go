package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestOverlay_1(t *testing.T) {
	type args struct {
		src ImageSource
		x   any
		y   any
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
			if got := f.Overlay(tt.args.src, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filters.Overlay() = %v, want %v", got, tt.want)
			}
		})
	}
}
