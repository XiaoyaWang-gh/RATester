package media

import (
	"fmt"
	"testing"
)

func TestIsHTML_1(t *testing.T) {
	tests := []struct {
		name string
		m    Type
		want bool
	}{
		{
			name: "HTML type",
			m:    Type{SubType: Builtin.HTMLType.SubType},
			want: true,
		},
		{
			name: "Non-HTML type",
			m:    Type{SubType: "text/plain"},
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

			if got := tt.m.IsHTML(); got != tt.want {
				t.Errorf("IsHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
