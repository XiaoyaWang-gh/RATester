package herrors

import (
	"fmt"
	"strings"
	"testing"
)

func TestContainsMatcher_1(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want func(m LineMatcher) int
	}{
		{
			name: "Test case 1",
			args: args{text: "test"},
			want: func(m LineMatcher) int {
				if idx := strings.Index(m.Line, "test"); idx != -1 {
					return idx + 1
				}
				return -1
			},
		},
		{
			name: "Test case 2",
			args: args{text: "example"},
			want: func(m LineMatcher) int {
				if idx := strings.Index(m.Line, "example"); idx != -1 {
					return idx + 1
				}
				return -1
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := ContainsMatcher(tt.args.text)
			matcher := LineMatcher{
				Line: "This is a test line",
			}
			if got(matcher) != tt.want(matcher) {
				t.Errorf("ContainsMatcher() = %v, want %v", got(matcher), tt.want(matcher))
			}
		})
	}
}
