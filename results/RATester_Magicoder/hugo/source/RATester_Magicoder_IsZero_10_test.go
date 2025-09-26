package source

import (
	"fmt"
	"testing"
)

func TestIsZero_10(t *testing.T) {
	tests := []struct {
		name string
		g    GitInfo
		want bool
	}{
		{
			name: "empty GitInfo",
			g:    GitInfo{},
			want: true,
		},
		{
			name: "non-empty GitInfo",
			g:    GitInfo{Hash: "abc123"},
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

			if got := tt.g.IsZero(); got != tt.want {
				t.Errorf("GitInfo.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
