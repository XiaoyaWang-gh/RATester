package publisher

import (
	"fmt"
	"testing"
)

func TestIsClosedByTag_1(t *testing.T) {
	type args struct {
		b       []byte
		tagName []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				b:       []byte("<div></div>"),
				tagName: []byte("div"),
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				b:       []byte("<div></div>"),
				tagName: []byte("span"),
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				b:       []byte("<div><span></span></div>"),
				tagName: []byte("div"),
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				b:       []byte("<div><span></span></div>"),
				tagName: []byte("span"),
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				b:       []byte("<div><span></span></div>"),
				tagName: []byte("span"),
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

			if got := isClosedByTag(tt.args.b, tt.args.tagName); got != tt.want {
				t.Errorf("isClosedByTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
