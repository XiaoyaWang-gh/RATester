package types

import (
	"fmt"
	"testing"
)

func TestContains_1(t *testing.T) {
	tests := []struct {
		name           string
		httpCodeRanges HTTPCodeRanges
		statusCode     int
		want           bool
	}{
		{
			name: "contains",
			httpCodeRanges: HTTPCodeRanges{
				{200, 299},
				{300, 399},
			},
			statusCode: 250,
			want:       true,
		},
		{
			name: "does not contain",
			httpCodeRanges: HTTPCodeRanges{
				{200, 299},
				{300, 399},
			},
			statusCode: 100,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.httpCodeRanges.Contains(tt.statusCode); got != tt.want {
				t.Errorf("HTTPCodeRanges.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
