package paths

import (
	"fmt"
	"testing"
)

func TestToSlashTrimLeading_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test case 1",
			args: "/path/to/file",
			want: "path/to/file",
		},
		{
			name: "Test case 2",
			args: "//path//to//file",
			want: "path/to/file",
		},
		{
			name: "Test case 3",
			args: "path/to/file",
			want: "path/to/file",
		},
		{
			name: "Test case 4",
			args: "path/to/file/",
			want: "path/to/file",
		},
		{
			name: "Test case 5",
			args: "/path/to/file/",
			want: "path/to/file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ToSlashTrimLeading(tt.args); got != tt.want {
				t.Errorf("ToSlashTrimLeading() = %v, want %v", got, tt.want)
			}
		})
	}
}
