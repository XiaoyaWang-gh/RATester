package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultColors_1(t *testing.T) {
	t.Parallel()

	type args struct {
		colors Colors
	}

	tests := []struct {
		name string
		args args
		want Colors
	}{
		{
			name: "Test with empty colors",
			args: args{
				colors: Colors{},
			},
			want: DefaultColors,
		},
		{
			name: "Test with some colors",
			args: args{
				colors: Colors{
					Black: "\u001b[30m",
					Red:   "\u001b[31m",
				},
			},
			want: Colors{
				Black:   "\u001b[30m",
				Red:     "\u001b[31m",
				Green:   DefaultColors.Green,
				Yellow:  DefaultColors.Yellow,
				Blue:    DefaultColors.Blue,
				Magenta: DefaultColors.Magenta,
				Cyan:    DefaultColors.Cyan,
				White:   DefaultColors.White,
				Reset:   DefaultColors.Reset,
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

			if got := defaultColors(tt.args.colors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
