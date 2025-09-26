package env

import (
	"fmt"
	"os"
	"testing"
)

func TestGetGOPATH_1(t *testing.T) {

	tests := []struct {
		name     string
		envVar   string
		envValue string
		want     string
	}{
		{
			name:     "GOPATH is set",
			envVar:   "GOPATH",
			envValue: "/usr/local/go",
			want:     "/usr/local/go",
		},
		{
			name:     "GOPATH is not set",
			envVar:   "GOPATH",
			envValue: "",
			want:     "/usr/local/go",
		},
		{
			name:     "GOPATH is set to empty string",
			envVar:   "GOPATH",
			envValue: " ",
			want:     "/usr/local/go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Setenv(tt.envVar, tt.envValue)
			defer os.Unsetenv(tt.envVar)

			if got := GetGOPATH(); got != tt.want {
				t.Errorf("GetGOPATH() = %v, want %v", got, tt.want)
			}
		})
	}
}
