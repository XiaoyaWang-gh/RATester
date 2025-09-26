package paths

import (
	"fmt"
	"testing"
)

func TestBase_1(t *testing.T) {
	bridge := pathBridge{}

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "simple path",
			in:   "/a/b/c",
			want: "c",
		},
		{
			name: "relative path",
			in:   "a/b/c",
			want: "c",
		},
		{
			name: "no path",
			in:   "",
			want: ".",
		},
		{
			name: "root path",
			in:   "/",
			want: "/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := bridge.Base(tt.in); got != tt.want {
				t.Errorf("pathBridge.Base() = %v, want %v", got, tt.want)
			}
		})
	}
}
