package publisher

import (
	"fmt"
	"testing"
)

func TestparseStartTag_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"Test case 1", "<div>", "div"},
		{"Test case 2", "<div class=\"test\">", "div"},
		{"Test case 3", "<div class=\"test\" />", "div"},
		{"Test case 4", "<div/>", "div"},
		{"Test case 5", "<div class=\"test\" />", "div"},
		{"Test case 6", "<div class=\"test\" />", "div"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := parseStartTag(tt.args); got != tt.want {
				t.Errorf("parseStartTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
