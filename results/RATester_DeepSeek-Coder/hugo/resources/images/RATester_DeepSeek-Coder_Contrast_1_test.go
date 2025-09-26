package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestContrast_1(t *testing.T) {
	t.Parallel()

	type args struct {
		percentage any
	}

	tests := []struct {
		name string
		args args
		want gift.Filter
	}{
		{
			name: "Test with valid percentage",
			args: args{percentage: 0.5},
			want: gift.Contrast(0.5),
		},
		{
			name: "Test with invalid percentage",
			args: args{percentage: "invalid"},
			want: gift.Contrast(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &Filters{}
			if got := f.Contrast(tt.args.percentage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Contrast() = %v, want %v", got, tt.want)
			}
		})
	}
}
