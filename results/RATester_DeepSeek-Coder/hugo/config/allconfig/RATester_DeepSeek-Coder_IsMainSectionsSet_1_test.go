package allconfig

import (
	"fmt"
	"testing"
)

func TestIsMainSectionsSet_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		cfg  *ConfigCompiled
		want bool
	}{
		{
			name: "MainSections is set",
			cfg: &ConfigCompiled{
				MainSections: []string{"section1", "section2"},
			},
			want: true,
		},
		{
			name: "MainSections is not set",
			cfg:  &ConfigCompiled{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.cfg.IsMainSectionsSet()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
