package paths

import (
	"fmt"
	"testing"
)

func TestContainer_1(t *testing.T) {
	tests := []struct {
		name string
		p    *Path
		want string
	}{
		{
			name: "Test case 1",
			p: &Path{
				s:                "test/container/path",
				posContainerLow:  0,
				posContainerHigh: 10,
			},
			want: "test/container",
		},
		{
			name: "Test case 2",
			p: &Path{
				s:                "test/path",
				posContainerLow:  -1,
				posContainerHigh: 0,
			},
			want: "",
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

			if got := tt.p.Container(); got != tt.want {
				t.Errorf("Container() = %v, want %v", got, tt.want)
			}
		})
	}
}
