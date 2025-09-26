package types

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	tests := []struct {
		name string
		prs  PortsRangeSlice
		want string
	}{
		{
			name: "empty slice",
			prs:  PortsRangeSlice{},
			want: "",
		},
		{
			name: "single port",
			prs: PortsRangeSlice{
				{Single: 8080},
			},
			want: "8080",
		},
		{
			name: "range of ports",
			prs: PortsRangeSlice{
				{Start: 8080, End: 8085},
			},
			want: "8080-8085",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.prs.String(); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
