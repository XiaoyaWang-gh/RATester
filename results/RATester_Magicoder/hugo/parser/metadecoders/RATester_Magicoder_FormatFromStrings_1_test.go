package metadecoders

import (
	"fmt"
	"testing"
)

func TestFormatFromStrings_1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want Format
	}{
		{
			name: "Test case 1",
			args: []string{"", "test"},
			want: "test",
		},
		{
			name: "Test case 2",
			args: []string{"", ""},
			want: "",
		},
		{
			name: "Test case 3",
			args: []string{"test", "test2"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FormatFromStrings(tt.args...); got != tt.want {
				t.Errorf("FormatFromStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
