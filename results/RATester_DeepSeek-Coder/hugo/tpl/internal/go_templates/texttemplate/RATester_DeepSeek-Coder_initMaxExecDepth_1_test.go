package template

import (
	"fmt"
	"testing"
)

func TestInitMaxExecDepth_1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Test case 1",
			want: 100000,
		},
		{
			name: "Test case 2",
			want: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := initMaxExecDepth()
			if got != tt.want {
				t.Errorf("initMaxExecDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
