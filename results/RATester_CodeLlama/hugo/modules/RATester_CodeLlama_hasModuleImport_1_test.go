package modules

import (
	"fmt"
	"testing"
)

func TestHasModuleImport_1(t *testing.T) {
	tests := []struct {
		name string
		c    Config
		want bool
	}{
		{
			name: "empty",
			c:    Config{},
			want: false,
		},
		{
			name: "has",
			c: Config{
				Imports: []Import{
					{
						Path: "github.com/bep/my-theme",
					},
				},
			},
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

			if got := tt.c.hasModuleImport(); got != tt.want {
				t.Errorf("Config.hasModuleImport() = %v, want %v", got, tt.want)
			}
		})
	}
}
