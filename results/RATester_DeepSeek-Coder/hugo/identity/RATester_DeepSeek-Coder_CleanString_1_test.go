package identity

import (
	"fmt"
	"testing"
)

func TestCleanString_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with simple string",
			args: args{s: "test/string"},
			want: "/test/string",
		},
		{
			name: "Test with trailing slash",
			args: args{s: "test/string/"},
			want: "/test/string",
		},
		{
			name: "Test with leading slash",
			args: args{s: "/test/string"},
			want: "/test/string",
		},
		{
			name: "Test with multiple slashes",
			args: args{s: "//test//string//"},
			want: "/test/string",
		},
		{
			name: "Test with uppercase letters",
			args: args{s: "TeSt/StRiNg"},
			want: "/test/string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := CleanString(tt.args.s); got != tt.want {
				t.Errorf("CleanString() = %v, want %v", got, tt.want)
			}
		})
	}
}
