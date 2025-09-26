package deploy

import (
	"fmt"
	"testing"
)

func TeststripIndexHTML_1(t *testing.T) {
	tests := []struct {
		name      string
		slashpath string
		want      string
	}{
		{
			name:      "Test case 1",
			slashpath: "/index.html",
			want:      "/",
		},
		{
			name:      "Test case 2",
			slashpath: "/test/index.html",
			want:      "/test/",
		},
		{
			name:      "Test case 3",
			slashpath: "/test",
			want:      "/test",
		},
		{
			name:      "Test case 4",
			slashpath: "/test/",
			want:      "/test/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := stripIndexHTML(tt.slashpath); got != tt.want {
				t.Errorf("stripIndexHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
