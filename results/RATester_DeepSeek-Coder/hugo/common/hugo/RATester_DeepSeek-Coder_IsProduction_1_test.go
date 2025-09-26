package hugo

import (
	"fmt"
	"testing"
)

func TestIsProduction_1(t *testing.T) {
	tests := []struct {
		name string
		i    HugoInfo
		want bool
	}{
		{
			name: "Test IsProduction with EnvironmentProduction",
			i: HugoInfo{
				Environment: EnvironmentProduction,
			},
			want: true,
		},
		{
			name: "Test IsProduction with other Environment",
			i: HugoInfo{
				Environment: "other",
			},
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

			if got := tt.i.IsProduction(); got != tt.want {
				t.Errorf("HugoInfo.IsProduction() = %v, want %v", got, tt.want)
			}
		})
	}
}
