package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
	"github.com/spf13/cast"
)

func TestSepia_1(t *testing.T) {
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
			args: args{
				percentage: 0.5,
			},
			want: filter{
				Options: newFilterOpts(0.5),
				Filter:  gift.Sepia(cast.ToFloat32(0.5)),
			},
		},
		{
			name: "Test with invalid percentage",
			args: args{
				percentage: "invalid",
			},
			want: filter{
				Options: newFilterOpts("invalid"),
				Filter:  gift.Sepia(cast.ToFloat32("invalid")),
			},
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
			if got := f.Sepia(tt.args.percentage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sepia() = %v, want %v", got, tt.want)
			}
		})
	}
}
