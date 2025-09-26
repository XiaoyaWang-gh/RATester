package hugo

import (
	"fmt"
	"testing"
)

func TestcompareFloatWithVersion_1(t *testing.T) {
	type args struct {
		v1 float64
		v2 Version
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				v1: 1.23,
				v2: Version{
					Major: 1,
					Minor: 23,
				},
			},
			want: 0,
		},
		{
			name: "Test case 2",
			args: args{
				v1: 2.34,
				v2: Version{
					Major: 1,
					Minor: 23,
				},
			},
			want: 1,
		},
		{
			name: "Test case 3",
			args: args{
				v1: 0.23,
				v2: Version{
					Major: 1,
					Minor: 23,
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := compareFloatWithVersion(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("compareFloatWithVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
