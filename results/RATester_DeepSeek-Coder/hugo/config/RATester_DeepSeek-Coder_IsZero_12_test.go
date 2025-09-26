package config

import (
	"fmt"
	"testing"
)

func TestIsZero_12(t *testing.T) {
	tests := []struct {
		name string
		r    Redirect
		want bool
	}{
		{
			name: "Test case 1",
			r:    Redirect{From: "", To: "test", Status: 301, Force: true},
			want: true,
		},
		{
			name: "Test case 2",
			r:    Redirect{From: "test", To: "test", Status: 301, Force: true},
			want: false,
		},
		{
			name: "Test case 3",
			r:    Redirect{From: "test", To: "", Status: 301, Force: true},
			want: false,
		},
		{
			name: "Test case 4",
			r:    Redirect{From: "test", To: "test", Status: 0, Force: true},
			want: false,
		},
		{
			name: "Test case 5",
			r:    Redirect{From: "test", To: "test", Status: 301, Force: false},
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

			if got := tt.r.IsZero(); got != tt.want {
				t.Errorf("Redirect.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
