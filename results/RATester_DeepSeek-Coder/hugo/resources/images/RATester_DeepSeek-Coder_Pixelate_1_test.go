package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestPixelate_1(t *testing.T) {
	t.Parallel()

	type args struct {
		size any
	}

	tests := []struct {
		name string
		args args
		want gift.Filter
	}{
		{
			name: "Test Case 1",
			args: args{
				size: 10,
			},
			want: gift.Pixelate(10),
		},
		{
			name: "Test Case 2",
			args: args{
				size: 20,
			},
			want: gift.Pixelate(20),
		},
		{
			name: "Test Case 3",
			args: args{
				size: 30,
			},
			want: gift.Pixelate(30),
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
			if got := f.Pixelate(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pixelate() = %v, want %v", got, tt.want)
			}
		})
	}
}
