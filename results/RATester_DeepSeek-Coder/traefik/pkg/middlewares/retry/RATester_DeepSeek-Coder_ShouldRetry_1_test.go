package retry

import (
	"fmt"
	"testing"
)

func TestShouldRetry_1(t *testing.T) {
	tests := []struct {
		name string
		r    *responseWriter
		want bool
	}{
		{
			name: "ShouldRetry returns true when shouldRetry is true",
			r: &responseWriter{
				shouldRetry: true,
			},
			want: true,
		},
		{
			name: "ShouldRetry returns false when shouldRetry is false",
			r: &responseWriter{
				shouldRetry: false,
			},
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

			if got := tt.r.ShouldRetry(); got != tt.want {
				t.Errorf("responseWriter.ShouldRetry() = %v, want %v", got, tt.want)
			}
		})
	}
}
