package hugo

import (
	"fmt"
	"testing"
)

func TestIsDartSassV2_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Test case 1",
			want: true,
		},
		{
			name: "Test case 2",
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

			if got := IsDartSassV2(); got != tt.want {
				t.Errorf("IsDartSassV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
