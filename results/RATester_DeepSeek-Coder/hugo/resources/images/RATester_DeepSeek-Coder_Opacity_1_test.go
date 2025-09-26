package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestOpacity_1(t *testing.T) {
	type args struct {
		opacity any
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
			if got := f.Opacity(tt.args.opacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Opacity() = %v, want %v", got, tt.want)
			}
		})
	}
}
