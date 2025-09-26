package gin

import (
	"fmt"
	"testing"
)

func TestIsASCII_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test with ASCII string",
			args: args{s: "Hello, World!"},
			want: true,
		},
		{
			name: "Test with non-ASCII string",
			args: args{s: "Héllo, World!"},
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

			if got := isASCII(tt.args.s); got != tt.want {
				t.Errorf("isASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
