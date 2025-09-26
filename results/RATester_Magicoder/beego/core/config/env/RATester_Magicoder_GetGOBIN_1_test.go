package env

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"testing"
)

func TestGetGOBIN_1(t *testing.T) {

	tests := []struct {
		name     string
		envVar   string
		envValue string
		want     string
	}{
		{
			name:     "Test case 1: GOBIN is set",
			envVar:   "GOBIN",
			envValue: "/usr/local/bin",
			want:     "/usr/local/bin",
		},
		{
			name:     "Test case 2: GOBIN is not set",
			envVar:   "",
			envValue: "",
			want:     filepath.Join(build.Default.GOPATH, "bin"),
		},
		{
			name:     "Test case 3: GOBIN is empty",
			envVar:   "GOBIN",
			envValue: "",
			want:     filepath.Join(build.Default.GOPATH, "bin"),
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

			if got := GetGOBIN(); got != tt.want {
				t.Errorf("GetGOBIN() = %v, want %v", got, tt.want)
			}
		})
	}
}
