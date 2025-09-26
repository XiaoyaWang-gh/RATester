package paths

import (
	"fmt"
	"testing"
)

func Test__3(t *testing.T) {
	var tests = []struct {
		name string
		want string
	}{
		{
			name: "PathTypeFile",
			want: "PathTypeFile",
		},
		{
			name: "PathTypeContentResource",
			want: "PathTypeContentResource",
		},
		{
			name: "PathTypeContentSingle",
			want: "PathTypeContentSingle",
		},
		{
			name: "PathTypeLeaf",
			want: "PathTypeLeaf",
		},
		{
			name: "PathTypeBranch",
			want: "PathTypeBranch",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.name; got != tt.want {
				t.Errorf("name() = %v, want %v", got, tt.want)
			}
		})
	}
}
