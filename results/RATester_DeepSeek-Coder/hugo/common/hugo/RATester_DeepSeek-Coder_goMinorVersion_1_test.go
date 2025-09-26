package hugo

import (
	"fmt"
	"testing"
)

func TestGoMinorVersion_1(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    int
	}{
		{"Test with valid version", "go1.16", 16},
		{"Test with devel version", "devel", 9999},
		{"Test with invalid version", "go1.x", 0},
		{"Test with no version", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := goMinorVersion(tt.version); got != tt.want {
				t.Errorf("goMinorVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
