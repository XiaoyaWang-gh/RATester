package media

import (
	"fmt"
	"testing"
)

func TestIsTextSuffix_1(t *testing.T) {
	tests := []struct {
		name   string
		types  Types
		suffix string
		want   bool
	}{
		{
			name: "test1",
			types: Types{
				{
					Type: "text/plain",
				},
				{
					Type: "application/json",
				},
			},
			suffix: "txt",
			want:   true,
		},
		{
			name: "test2",
			types: Types{
				{
					Type: "text/plain",
				},
				{
					Type: "application/json",
				},
			},
			suffix: "json",
			want:   false,
		},
		{
			name: "test3",
			types: Types{
				{
					Type: "text/plain",
				},
				{
					Type: "application/json",
				},
			},
			suffix: "xml",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.types.IsTextSuffix(tt.suffix); got != tt.want {
				t.Errorf("IsTextSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
