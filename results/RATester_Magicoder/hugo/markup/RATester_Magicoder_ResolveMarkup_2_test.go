package markup

import (
	"fmt"
	"testing"
)

func TestResolveMarkup_2(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test case 1",
			args: "goldmark",
			want: "markdown",
		},
		{
			name: "Test case 2",
			args: "asciidocext",
			want: "asciidoc",
		},
		{
			name: "Test case 3",
			args: "unknown",
			want: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ResolveMarkup(tt.args); got != tt.want {
				t.Errorf("ResolveMarkup() = %v, want %v", got, tt.want)
			}
		})
	}
}
