package traefik

import (
	"fmt"
	"testing"
	"time"
)

func TestThrottleDuration_2(t *testing.T) {
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

			i := Provider{}
			if got := i.ThrottleDuration(); got != tt.want {
				t.Errorf("ThrottleDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
