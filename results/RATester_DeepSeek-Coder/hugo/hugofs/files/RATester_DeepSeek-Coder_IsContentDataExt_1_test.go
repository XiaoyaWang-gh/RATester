package files

import (
	"fmt"
	"testing"
)

func TestIsContentDataExt_1(t *testing.T) {
	tests := []struct {
		name string
		ext  string
		want bool
	}{
		{
			name: "Test .tmpl",
			ext:  ".tmpl",
			want: true,
		},
		{
			name: "Test .html",
			ext:  ".html",
			want: false,
		},
		{
			name: "Test .txt",
			ext:  ".txt",
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

			if got := IsContentDataExt(tt.ext); got != tt.want {
				t.Errorf("IsContentDataExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
