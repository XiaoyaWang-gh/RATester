package metrics

import (
	"fmt"
	"testing"
)

func TestHasRouter_1(t *testing.T) {
	d := &dynamicConfig{
		routers: map[string]bool{
			"router1": true,
			"router2": true,
		},
	}

	tests := []struct {
		name   string
		router string
		want   bool
	}{
		{
			name:   "Existing router",
			router: "router1",
			want:   true,
		},
		{
			name:   "Non-existing router",
			router: "router3",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := d.hasRouter(tt.router); got != tt.want {
				t.Errorf("hasRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
