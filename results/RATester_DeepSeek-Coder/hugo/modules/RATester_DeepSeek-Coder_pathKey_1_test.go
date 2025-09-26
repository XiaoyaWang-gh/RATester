package modules

import (
	"fmt"
	"testing"
)

func TestPathKey_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test case 1",
			args: "github.com/user/repo@v1.0.0",
			want: "github.com/user/repo",
		},
		{
			name: "Test case 2",
			args: "github.com/user/repo2@v2.0.0",
			want: "github.com/user/repo2",
		},
		{
			name: "Test case 3",
			args: "github.com/user/repo3@v3.0.0",
			want: "github.com/user/repo3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := pathKey(tt.args); got != tt.want {
				t.Errorf("pathKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
