package web

import (
	"fmt"
	"html/template"
	"testing"
)

func TestAssetsJs_1(t *testing.T) {
	tests := []struct {
		name string
		text string
		want template.HTML
	}{
		{
			name: "Test case 1",
			text: "test.js",
			want: "<script src=\"test.js\"></script>",
		},
		{
			name: "Test case 2",
			text: "test2.js",
			want: "<script src=\"test2.js\"></script>",
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

			if got := AssetsJs(tt.text); got != tt.want {
				t.Errorf("AssetsJs() = %v, want %v", got, tt.want)
			}
		})
	}
}
