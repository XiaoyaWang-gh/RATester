package paths

import (
	"fmt"
	"testing"
)

func TestString_9(t *testing.T) {
	tests := []struct {
		name string
		i    PathType
		want string
	}{
		{
			name: "Test PathType 0",
			i:    0,
			want: "PathType(0)",
		},
		{
			name: "Test PathType 1",
			i:    1,
			want: "PathType(1)",
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

			if got := tt.i.String(); got != tt.want {
				t.Errorf("PathType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
