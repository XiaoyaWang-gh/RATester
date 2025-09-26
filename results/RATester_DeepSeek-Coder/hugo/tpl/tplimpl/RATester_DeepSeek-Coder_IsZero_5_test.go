package tplimpl

import (
	"fmt"
	"testing"
)

func TestIsZero_5(t *testing.T) {
	tests := []struct {
		name     string
		template templateInfo
		want     bool
	}{
		{
			name: "Test with non-empty name",
			template: templateInfo{
				name: "test",
			},
			want: false,
		},
		{
			name: "Test with empty name",
			template: templateInfo{
				name: "",
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

			if got := tt.template.IsZero(); got != tt.want {
				t.Errorf("templateInfo.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
