package gin

import (
	"fmt"
	"testing"
)

func TestcalculateAbsolutePath_1(t *testing.T) {
	group := &RouterGroup{
		basePath: "/base",
	}

	type args struct {
		relativePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				relativePath: "/relative",
			},
			want: "/base/relative",
		},
		{
			name: "Test case 2",
			args: args{
				relativePath: "/sub/relative",
			},
			want: "/base/sub/relative",
		},
		{
			name: "Test case 3",
			args: args{
				relativePath: "/",
			},
			want: "/base/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := group.calculateAbsolutePath(tt.args.relativePath); got != tt.want {
				t.Errorf("calculateAbsolutePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
