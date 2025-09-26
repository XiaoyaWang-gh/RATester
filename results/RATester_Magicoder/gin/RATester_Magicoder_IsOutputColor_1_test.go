package gin

import (
	"fmt"
	"testing"
)

func TestIsOutputColor_1(t *testing.T) {
	tests := []struct {
		name string
		p    *LogFormatterParams
		want bool
	}{
		{
			name: "Test case 1",
			p: &LogFormatterParams{
				isTerm: true,
			},
			want: true,
		},
		{
			name: "Test case 2",
			p: &LogFormatterParams{
				isTerm: false,
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

			if got := tt.p.IsOutputColor(); got != tt.want {
				t.Errorf("LogFormatterParams.IsOutputColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
