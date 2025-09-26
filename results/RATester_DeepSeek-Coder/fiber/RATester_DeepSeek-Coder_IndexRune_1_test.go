package fiber

import (
	"fmt"
	"testing"
)

func TestIndexRune_1(t *testing.T) {
	type args struct {
		str    string
		needle int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				str:    "hello",
				needle: 'e',
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				str:    "hello",
				needle: 'a',
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				str:    "日本語",
				needle: '語',
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				str:    "日本語",
				needle: 'a',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IndexRune(tt.args.str, tt.args.needle); got != tt.want {
				t.Errorf("IndexRune() = %v, want %v", got, tt.want)
			}
		})
	}
}
