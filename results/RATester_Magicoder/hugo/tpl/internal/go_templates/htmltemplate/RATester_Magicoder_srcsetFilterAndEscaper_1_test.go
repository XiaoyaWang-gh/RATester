package template

import (
	"fmt"
	"testing"
)

func TestsrcsetFilterAndEscaper_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test Case 1",
			args: []any{"test1", "test2"},
			want: "test1,test2",
		},
		{
			name: "Test Case 2",
			args: []any{"test3", "test4"},
			want: "test3,test4",
		},
		// Add more test cases as needed
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
