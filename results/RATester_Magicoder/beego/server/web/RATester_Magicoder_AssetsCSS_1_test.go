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
			name: "Test case 1",
			args: args{
				text: "test.css",
			},
			want: "<link href=\"test.css\" rel=\"stylesheet\" />",
		},
		{
			name: "Test case 2",
			args: args{
				text: "test2.css",
			},
			want: "<link href=\"test2.css\" rel=\"stylesheet\" />",
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
