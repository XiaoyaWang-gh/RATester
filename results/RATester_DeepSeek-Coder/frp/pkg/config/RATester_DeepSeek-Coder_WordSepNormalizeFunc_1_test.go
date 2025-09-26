package config

import (
	"fmt"
	"testing"

	"github.com/spf13/pflag"
)

func TestWordSepNormalizeFunc_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want pflag.NormalizedName
	}{
		{
			name: "test_case_1",
			want: pflag.NormalizedName("test-case-1"),
		},
		{
			name: "test_case_2",
			want: pflag.NormalizedName("test-case-2"),
		},
		{
			name: "test_case_3",
			want: pflag.NormalizedName("test-case-3"),
		},
	}

	for _, tt := range tests {
		tt := tt // NOTE: https://golang.org/doc/faq#closures_and_goroutines

		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			if got := WordSepNormalizeFunc(nil, tt.name); got != tt.want {
				t.Errorf("WordSepNormalizeFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
