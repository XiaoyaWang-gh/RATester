package paths

import (
	"fmt"
	"testing"
)

func TestAddContextRoot_1(t *testing.T) {
	type args struct {
		baseURL      string
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
				baseURL:      "http://localhost:8080/api/v1",
				relativePath: "users",
			},
			want: "http://localhost:8080/api/v1/users",
		},
		{
			name: "Test case 2",
			args: args{
				baseURL:      "http://localhost:8080/api/v1/",
				relativePath: "users/",
			},
			want: "http://localhost:8080/api/v1/users/",
		},
		{
			name: "Test case 3",
			args: args{
				baseURL:      "http://localhost:8080/api/v1",
				relativePath: "/users",
			},
			want: "http://localhost:8080/api/v1/users",
		},
		{
			name: "Test case 4",
			args: args{
				baseURL:      "http://localhost:8080/api/v1/",
				relativePath: "/users/",
			},
			want: "http://localhost:8080/api/v1/users/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AddContextRoot(tt.args.baseURL, tt.args.relativePath); got != tt.want {
				t.Errorf("AddContextRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
