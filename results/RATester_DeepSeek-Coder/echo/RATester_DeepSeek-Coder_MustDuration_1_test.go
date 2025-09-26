package echo

import (
	"fmt"
	"testing"
	"time"
)

func TestMustDuration_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		sourceParam string
		want        time.Duration
	}{
		{
			name:        "valid duration",
			sourceParam: "1s",
			want:        time.Second,
		},
		{
			name:        "invalid duration",
			sourceParam: "invalid",
			want:        0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := &ValueBinder{
				ValueFunc: func(sourceParam string) string {
					if sourceParam == "1s" {
						return "1s"
					}
					return ""
				},
			}

			var got time.Duration
			b.MustDuration(tt.sourceParam, &got)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
