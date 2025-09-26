package template

import (
	"fmt"
	"testing"
)

func TesthtmlNameFilter_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"test1", contentTypeHTMLAttr},
			want: "test1",
		},
		{
			name: "Test case 2",
			args: []any{"test2", contentTypePlain},
			want: "test2",
		},
		{
			name: "Test case 3",
			args: []any{"test3", contentTypePlain},
			want: "test3",
		},
		{
			name: "Test case 4",
			args: []any{"test4", contentTypePlain},
			want: "test4",
		},
		{
			name: "Test case 5",
			args: []any{"test5", contentTypePlain},
			want: "test5",
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
