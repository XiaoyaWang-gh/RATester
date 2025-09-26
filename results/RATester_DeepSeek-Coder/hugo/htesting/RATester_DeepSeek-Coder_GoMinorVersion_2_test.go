package htesting

import (
	"fmt"
	"testing"
)

func TestGoMinorVersion_2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Test case 1",
			want: 1,
		},
		{
			name: "Test case 2",
			want: 2,
		},
		{
			name: "Test case 3",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GoMinorVersion(); got != tt.want {
				t.Errorf("GoMinorVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
