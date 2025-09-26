package helpers

import (
	"fmt"
	"testing"
)

func TestgetBaseURLRoot_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with absolute path",
			args: args{
				path: "/absolute/path",
			},
			want: "WithoutPath",
		},
		{
			name: "Test with relative path",
			args: args{
				path: "relative/path",
			},
			want: "WithPath",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &PathSpec{}
			if got := p.getBaseURLRoot(tt.args.path); got != tt.want {
				t.Errorf("getBaseURLRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
