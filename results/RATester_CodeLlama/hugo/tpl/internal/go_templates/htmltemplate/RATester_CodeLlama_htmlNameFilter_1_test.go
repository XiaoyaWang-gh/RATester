package template

import (
	"fmt"
	"testing"
)

func TestHtmlNameFilter_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "test1",
			args: []any{"test1"},
			want: "test1",
		},
		{
			name: "test2",
			args: []any{"test2"},
			want: "test2",
		},
		{
			name: "test3",
			args: []any{"test3"},
			want: "test3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := htmlNameFilter(tt.args...); got != tt.want {
				t.Errorf("htmlNameFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
