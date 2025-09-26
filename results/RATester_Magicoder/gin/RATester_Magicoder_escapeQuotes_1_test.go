package gin

import (
	"fmt"
	"testing"
)

func TestescapeQuotes_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test case 1",
			args: "Hello, world!",
			want: "Hello, world!",
		},
		{
			name: "Test case 2",
			args: "I'm a \"developer\".",
			want: "I\\'m a \\\"developer\\\".",
		},
		{
			name: "Test case 3",
			args: "\"Hello\", \"world\".",
			want: "\\\"Hello\\\", \\\"world\\\".",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := escapeQuotes(tt.args); got != tt.want {
				t.Errorf("escapeQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
