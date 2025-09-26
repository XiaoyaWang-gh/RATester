package livereload

import (
	"fmt"
	"testing"
)

func TestNavigateToPathForPort_1(t *testing.T) {
	type args struct {
		path string
		port int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				path: "/test",
				port: 8080,
			},
		},
		{
			name: "Test case 2",
			args: args{
				path: "/test2",
				port: 8081,
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			NavigateToPathForPort(tt.args.path, tt.args.port)
		})
	}
}
