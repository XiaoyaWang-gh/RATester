package internal

import (
	"fmt"
	"testing"
)

func TestPath_3(t *testing.T) {
	tests := []struct {
		name string
		d    ResourcePaths
		want string
	}{
		{
			name: "Test case 1",
			d: ResourcePaths{
				Dir:  "/path/to/dir",
				File: "data.json",
			},
			want: "/path/to/dir/data.json",
		},
		{
			name: "Test case 2",
			d: ResourcePaths{
				Dir:  "/path/to/dir",
				File: "",
			},
			want: "/path/to/dir",
		},
		{
			name: "Test case 3",
			d: ResourcePaths{
				Dir:  "",
				File: "data.json",
			},
			want: "data.json",
		},
		{
			name: "Test case 4",
			d: ResourcePaths{
				Dir:  "",
				File: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.Path(); got != tt.want {
				t.Errorf("Path() = %v, want %v", got, tt.want)
			}
		})
	}
}
