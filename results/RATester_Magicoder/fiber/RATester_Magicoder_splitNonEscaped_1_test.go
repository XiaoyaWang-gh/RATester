package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsplitNonEscaped_1(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				s:   "hello,world",
				sep: ",",
			},
			want: []string{"hello", "world"},
		},
		{
			name: "Test case 2",
			args: args{
				s:   "hello,world,",
				sep: ",",
			},
			want: []string{"hello", "world", ""},
		},
		{
			name: "Test case 3",
			args: args{
				s:   "hello,world,",
				sep: ",",
			},
			want: []string{"hello", "world", ""},
		},
		{
			name: "Test case 4",
			args: args{
				s:   "hello,world,",
				sep: ",",
			},
			want: []string{"hello", "world", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := splitNonEscaped(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitNonEscaped() = %v, want %v", got, tt.want)
			}
		})
	}
}
