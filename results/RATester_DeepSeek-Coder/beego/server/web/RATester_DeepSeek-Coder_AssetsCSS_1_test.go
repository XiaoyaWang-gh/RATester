package web

import (
	"fmt"
	"html/template"
	"testing"
)

func TestAssetsCSS_1(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want template.HTML
	}{
		{
			name: "Test with a valid CSS file path",
			args: args{text: "https://example.com/style.css"},
			want: "<link href=\"https://example.com/style.css\" rel=\"stylesheet\" />",
		},
		{
			name: "Test with a path containing a quote",
			args: args{text: "https://example.com/style\"s.css"},
			want: "<link href=\"https://example.com/style\\\"s.css\" rel=\"stylesheet\" />",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AssetsCSS(tt.args.text); got != tt.want {
				t.Errorf("AssetsCSS() = %v, want %v", got, tt.want)
			}
		})
	}
}
