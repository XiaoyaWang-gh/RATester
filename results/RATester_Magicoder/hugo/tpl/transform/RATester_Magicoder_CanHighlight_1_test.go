package transform

import (
	"fmt"
	"testing"
)

func TestCanHighlight_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		language string
		want     bool
	}{
		{
			name:     "valid language",
			language: "go",
			want:     true,
		},
		{
			name:     "invalid language",
			language: "invalid",
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

			if got := ns.CanHighlight(tt.language); got != tt.want {
				t.Errorf("CanHighlight() = %v, want %v", got, tt.want)
			}
		})
	}
}
