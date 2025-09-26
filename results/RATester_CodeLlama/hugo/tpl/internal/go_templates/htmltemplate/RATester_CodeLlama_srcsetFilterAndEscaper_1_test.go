package template

import (
	"fmt"
	"testing"
)

func TestSrcsetFilterAndEscaper_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "test_case_1",
			args: []any{
				"test_case_1",
			},
			want: "test_case_1",
		},
		{
			name: "test_case_2",
			args: []any{
				"test_case_2",
			},
			want: "test_case_2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := srcsetFilterAndEscaper(tt.args...); got != tt.want {
				t.Errorf("srcsetFilterAndEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
