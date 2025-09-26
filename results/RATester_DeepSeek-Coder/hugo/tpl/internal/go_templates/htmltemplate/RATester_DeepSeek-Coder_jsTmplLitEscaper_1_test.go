package template

import (
	"fmt"
	"testing"
)

func TestJsTmplLitEscaper_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"test1", "test2", "test3"},
			want: "test1test2test3",
		},
		{
			name: "Test case 2",
			args: []any{"<script>", "</script>"},
			want: "&lt;script&gt;&lt;/script&gt;",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := jsTmplLitEscaper(tt.args...); got != tt.want {
				t.Errorf("jsTmplLitEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
