package paths

import (
	"fmt"
	"testing"
)

func TestDir_1(t *testing.T) {
	bridge := pathBridge{}
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "TestDir/1",
			in:   "/a/b/c",
			want: "/a/b",
		},
		{
			name: "TestDir/2",
			in:   "/a/b/c/",
			want: "/a/b/c",
		},
		{
			name: "TestDir/3",
			in:   "a/b/c",
			want: "a/b",
		},
		{
			name: "TestDir/4",
			in:   "a/b/c/",
			want: "a/b/c",
		},
		{
			name: "TestDir/5",
			in:   "a",
			want: ".",
		},
		{
			name: "TestDir/6",
			in:   "/",
			want: "/",
		},
		{
			name: "TestDir/7",
			in:   "",
			want: ".",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := bridge.Dir(tt.in); got != tt.want {
				t.Errorf("Dir() = %v, want %v", got, tt.want)
			}
		})
	}
}
