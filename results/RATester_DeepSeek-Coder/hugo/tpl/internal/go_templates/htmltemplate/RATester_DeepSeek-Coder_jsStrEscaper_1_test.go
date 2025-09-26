package template

import (
	"fmt"
	"testing"
)

func TestJsStrEscaper_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"<script>", "</script>"},
			want: "\\u003Cscript\\u003E\\u003C\\/script\\u003E",
		},
		{
			name: "Test case 2",
			args: []any{"<", ">"},
			want: "\\u003C\\u003E",
		},
		{
			name: "Test case 3",
			args: []any{"\\", "\""},
			want: "\\u005C\\u0022",
		},
		{
			name: "Test case 4",
			args: []any{"\n", "\r"},
			want: "\\u000A\\u000D",
		},
		{
			name: "Test case 5",
			args: []any{"\t", " "},
			want: "\\u0009\\u0020",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := jsStrEscaper(tt.args...); got != tt.want {
				t.Errorf("jsStrEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
