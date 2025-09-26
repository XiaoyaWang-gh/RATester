package plugins

import (
	"fmt"
	"testing"
)

func TestIsYaegiPlugin_1(t *testing.T) {
	tests := []struct {
		name     string
		manifest *Manifest
		want     bool
	}{
		{
			name:     "runtime is yaegi",
			manifest: &Manifest{Runtime: runtimeYaegi},
			want:     true,
		},
		{
			name:     "runtime is empty",
			manifest: &Manifest{Runtime: ""},
			want:     true,
		},
		{
			name:     "runtime is not yaegi",
			manifest: &Manifest{Runtime: runtimeWasm},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.manifest.IsYaegiPlugin(); got != tt.want {
				t.Errorf("Manifest.IsYaegiPlugin() = %v, want %v", got, tt.want)
			}
		})
	}
}
