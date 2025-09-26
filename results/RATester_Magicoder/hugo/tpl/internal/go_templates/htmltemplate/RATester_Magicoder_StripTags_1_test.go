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
			name: "Test case 1",
			args: args{html: "<p>Hello, World!</p>"},
			want: "Hello, World!",
		},
		{
			name: "Test case 2",
			args: args{html: "<p>Hello, <b>World</b>!</p>"},
			want: "Hello, World!",
		},
		{
			name: "Test case 3",
			args: args{html: "<p>Hello, <b>World</b>!</p>"},
			want: "Hello, World!",
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
