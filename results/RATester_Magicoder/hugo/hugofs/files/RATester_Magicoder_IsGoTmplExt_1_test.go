package files

import (
	"fmt"
	"testing"
)

func TestIsGoTmplExt_1(t *testing.T) {
	tests := []struct {
		name string
		ext  string
		want bool
	}{
		{
			name: "Test case 1",
			ext:  "gotmpl",
			want: true,
		},
		{
			name: "Test case 2",
			ext:  "html",
			want: false,
		},
		{
			name: "Test case 3",
			ext:  "tmpl",
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

			if got := IsGoTmplExt(tt.ext); got != tt.want {
				t.Errorf("IsGoTmplExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
