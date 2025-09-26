package fiber

import (
	"fmt"
	"testing"
)

func TestGetGroupPath_1(t *testing.T) {
	type args struct {
		prefix string
		path   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				prefix: "/api",
				path:   "/v1",
			},
			want: "/api/v1",
		},
		{
			name: "Test Case 2",
			args: args{
				prefix: "/api/",
				path:   "v1",
			},
			want: "/api/v1",
		},
		{
			name: "Test Case 3",
			args: args{
				prefix: "/api",
				path:   "",
			},
			want: "/api",
		},
		{
			name: "Test Case 4",
			args: args{
				prefix: "/api/",
				path:   "",
			},
			want: "/api",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getGroupPath(tt.args.prefix, tt.args.path); got != tt.want {
				t.Errorf("getGroupPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
