package tailscale

import (
	"fmt"
	"testing"
	"time"
)

func TestThrottleDuration_3(t *testing.T) {
	tests := []struct {
		name string
		want time.Duration
	}{
		{
			name: "Test case 1",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{
				ResolverName: "test",
			}
			if got := p.ThrottleDuration(); got != tt.want {
				t.Errorf("Provider.ThrottleDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
