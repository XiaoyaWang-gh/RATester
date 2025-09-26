package modules

import (
	"fmt"
	"testing"
)

func TestIsProbablyModule_1(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "Test case 1",
			path: "github.com/user/repo",
			want: true,
		},
		{
			name: "Test case 2",
			path: "github.com/user/repo/cmd/subcmd",
			want: false,
		},
		{
			name: "Test case 3",
			path: "github.com/user/repo/...",
			want: true,
		},
		{
			name: "Test case 4",
			path: "github.com/user/repo/pkg",
			want: false,
		},
		{
			name: "Test case 5",
			path: "github.com/user/repo/v2",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isProbablyModule(tt.path); got != tt.want {
				t.Errorf("isProbablyModule() = %v, want %v", got, tt.want)
			}
		})
	}
}
