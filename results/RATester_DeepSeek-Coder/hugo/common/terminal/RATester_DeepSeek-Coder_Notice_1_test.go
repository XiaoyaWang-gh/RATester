package terminal

import (
	"fmt"
	"testing"
)

func TestNotice_1(t *testing.T) {
	testCases := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test Notice with string",
			arg:  "Hello, World",
			want: "\033[34mHello, World\033[0m",
		},
		{
			name: "Test Notice with empty string",
			arg:  "",
			want: "\033[34m\033[0m",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := Notice(tc.arg)
			if got != tc.want {
				t.Errorf("Notice(%q) = %q, want %q", tc.arg, got, tc.want)
			}
		})
	}
}
