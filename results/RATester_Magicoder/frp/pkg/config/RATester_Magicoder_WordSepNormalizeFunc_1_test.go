package config

import (
	"fmt"
	"testing"

	"github.com/spf13/pflag"
)

func TestWordSepNormalizeFunc_1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  pflag.NormalizedName
	}{
		{
			name:  "no_underscore",
			input: "test",
			want:  "test",
		},
		{
			name:  "with_underscore",
			input: "test_case",
			want:  "test-case",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WordSepNormalizeFunc(nil, tt.input); got != tt.want {
				t.Errorf("WordSepNormalizeFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
