package types

import (
	"fmt"
	"testing"
)

func TestContains_1(t *testing.T) {
	tests := []struct {
		name         string
		h            HTTPCodeRanges
		statusCode   int
		wantContains bool
	}{
		{
			name: "contains",
			h: HTTPCodeRanges{
				{200, 299},
				{300, 399},
			},
			statusCode:   250,
			wantContains: true,
		},
		{
			name: "does not contain",
			h: HTTPCodeRanges{
				{200, 299},
				{300, 399},
			},
			statusCode:   100,
			wantContains: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotContains := tt.h.Contains(tt.statusCode)
			if gotContains != tt.wantContains {
				t.Errorf("got %t, want %t", gotContains, tt.wantContains)
			}
		})
	}
}
