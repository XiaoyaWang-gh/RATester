package paths

import (
	"fmt"
	"testing"
)

func TestContainerDir_1(t *testing.T) {
	tests := []struct {
		name string
		p    *Path
		want string
	}{
		{
			name: "Test case 1",
			p: &Path{
				s:               "test/path",
				posContainerLow: -1,
			},
			want: "test",
		},
		{
			name: "Test case 2",
			p: &Path{
				s:               "test/path",
				posContainerLow: 0,
			},
			want: "",
		},
		{
			name: "Test case 3",
			p: &Path{
				s:               "test/path",
				posContainerLow: 5,
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.ContainerDir(); got != tt.want {
				t.Errorf("ContainerDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
