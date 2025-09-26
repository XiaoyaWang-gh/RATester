package common

import (
	"fmt"
	"testing"
)

func TestGetAppPath_1(t *testing.T) {
	type test struct {
		name string
		want string
	}

	tests := []test{
		{
			name: "Test case 1",
			want: "/path/to/app",
		},
		{
			name: "Test case 2",
			want: "/path/to/another/app",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetAppPath(); got != tt.want {
				t.Errorf("GetAppPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
