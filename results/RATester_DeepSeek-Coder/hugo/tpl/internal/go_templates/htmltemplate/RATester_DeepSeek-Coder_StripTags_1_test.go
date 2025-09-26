package template

import (
	"fmt"
	"testing"
)

func TestStripTags_1(t *testing.T) {
	type args struct {
		html string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with simple HTML",
			args: args{html: "<p>Hello, World</p>"},
			want: "Hello, World",
		},
		{
			name: "Test with nested HTML",
			args: args{html: "<div><p>Hello, <b>World</b></p></div>"},
			want: "Hello, World",
		},
		{
			name: "Test with multiple tags",
			args: args{html: "<p>Hello, <b>W</b>or<i>l</i>d</p>"},
			want: "Hello, Worlld",
		},
		{
			name: "Test with no tags",
			args: args{html: "Hello, World"},
			want: "Hello, World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := StripTags(tt.args.html); got != tt.want {
				t.Errorf("StripTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
