package sass

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types/css"
)

func TestIsTypedCSSValue_1(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want bool
	}{
		{
			name: "int",
			v:    1,
			want: true,
		},
		{
			name: "float64",
			v:    1.0,
			want: true,
		},
		{
			name: "string",
			v:    "red",
			want: true,
		},
		{
			name: "css.UnquotedString",
			v:    css.UnquotedString("red"),
			want: true,
		},
		{
			name: "string with color",
			v:    "#ff0000",
			want: true,
		},
		{
			name: "string with function",
			v:    "rgb(255, 0, 0)",
			want: true,
		},
		{
			name: "string with unit",
			v:    "10px",
			want: true,
		},
		{
			name: "unsupported type",
			v:    []int{1, 2, 3},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isTypedCSSValue(tt.v); got != tt.want {
				t.Errorf("isTypedCSSValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
