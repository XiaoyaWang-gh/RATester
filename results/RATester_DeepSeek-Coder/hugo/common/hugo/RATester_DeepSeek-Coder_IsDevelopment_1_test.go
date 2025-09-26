package hugo

import (
	"fmt"
	"testing"
)

func TestIsDevelopment_1(t *testing.T) {
	tests := []struct {
		name string
		i    HugoInfo
		want bool
	}{
		{
			name: "Test with development environment",
			i: HugoInfo{
				Environment: "development",
			},
			want: true,
		},
		{
			name: "Test with production environment",
			i: HugoInfo{
				Environment: "production",
			},
			want: false,
		},
		{
			name: "Test with empty environment",
			i: HugoInfo{
				Environment: "",
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

			if got := tt.i.IsDevelopment(); got != tt.want {
				t.Errorf("HugoInfo.IsDevelopment() = %v, want %v", got, tt.want)
			}
		})
	}
}
