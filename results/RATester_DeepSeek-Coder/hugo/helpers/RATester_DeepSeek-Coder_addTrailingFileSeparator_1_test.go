package helpers

import (
	"fmt"
	"testing"
)

func TestAddTrailingFileSeparator_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with trailing slash",
			args: args{s: "/path/to/dir/"},
			want: "/path/to/dir/",
		},
		{
			name: "Test without trailing slash",
			args: args{s: "/path/to/dir"},
			want: "/path/to/dir/",
		},
		{
			name: "Test with empty string",
			args: args{s: ""},
			want: "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := addTrailingFileSeparator(tt.args.s); got != tt.want {
				t.Errorf("addTrailingFileSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
