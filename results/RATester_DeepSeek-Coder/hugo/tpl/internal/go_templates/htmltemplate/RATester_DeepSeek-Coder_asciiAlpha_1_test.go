package template

import (
	"fmt"
	"testing"
)

func TestAsciiAlpha_1(t *testing.T) {
	testCases := []struct {
		name  string
		input byte
		want  bool
	}{
		{"Test lowercase a", 'a', true},
		{"Test lowercase z", 'z', true},
		{"Test uppercase A", 'A', true},
		{"Test uppercase Z", 'Z', true},
		{"Test lowercase before a", '`', false},
		{"Test lowercase after z", '{', false},
		{"Test uppercase before A", '@', false},
		{"Test uppercase after Z", '[', false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := asciiAlpha(tc.input)
			if got != tc.want {
				t.Errorf("asciiAlpha(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
