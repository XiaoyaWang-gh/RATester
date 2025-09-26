package paths

import (
	"fmt"
	"testing"
)

func TestExt_1(t *testing.T) {
	bridge := pathBridge{}

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "TestExt_1",
			in:   "/a/b/c/d.go",
			want: ".go",
		},
		{
			name: "TestExt_2",
			in:   "/a/b/c.d.go",
			want: ".go",
		},
		{
			name: "TestExt_3",
			in:   "/a/b/c",
			want: "",
		},
		{
			name: "TestExt_4",
			in:   "a/b/c",
			want: "",
		},
		{
			name: "TestExt_5",
			in:   "a.b.c",
			want: ".c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := bridge.Ext(tt.in); got != tt.want {
				t.Errorf("pathBridge.Ext() = %v, want %v", got, tt.want)
			}
		})
	}
}
