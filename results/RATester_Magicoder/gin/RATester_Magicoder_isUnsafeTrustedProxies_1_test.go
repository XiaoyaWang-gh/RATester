package gin

import (
	"fmt"
	"testing"
)

func TestisUnsafeTrustedProxies_1(t *testing.T) {
	engine := &Engine{
		trustedProxies: []string{"0.0.0.0", "::"},
	}

	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Test case 1",
			want: true,
		},
		{
			name: "Test case 2",
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

			if got := engine.isUnsafeTrustedProxies(); got != tt.want {
				t.Errorf("isUnsafeTrustedProxies() = %v, want %v", got, tt.want)
			}
		})
	}
}
