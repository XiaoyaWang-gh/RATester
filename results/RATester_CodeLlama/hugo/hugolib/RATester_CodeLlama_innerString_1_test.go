package hugolib

import (
	"fmt"
	"testing"
)

func TestInnerString_1(t *testing.T) {
	type args struct {
		s shortcode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s: shortcode{
					inner: []any{"test1"},
				},
			},
			want: "test1",
		},
		{
			name: "test2",
			args: args{
				s: shortcode{
					inner: []any{"test2"},
				},
			},
			want: "test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.s.innerString(); got != tt.want {
				t.Errorf("innerString() = %v, want %v", got, tt.want)
			}
		})
	}
}
