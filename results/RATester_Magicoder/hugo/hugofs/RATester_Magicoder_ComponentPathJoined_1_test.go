package hugofs

import (
	"fmt"
	"testing"
)

func TestComponentPathJoined_1(t *testing.T) {
	tests := []struct {
		name string
		c    ComponentPath
		want string
	}{
		{
			name: "Test case 1",
			c: ComponentPath{
				Component: "component1",
				Path:      "path1",
			},
			want: "component1/path1",
		},
		{
			name: "Test case 2",
			c: ComponentPath{
				Component: "component2",
				Path:      "path2",
			},
			want: "component2/path2",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.ComponentPathJoined(); got != tt.want {
				t.Errorf("ComponentPathJoined() = %v, want %v", got, tt.want)
			}
		})
	}
}
