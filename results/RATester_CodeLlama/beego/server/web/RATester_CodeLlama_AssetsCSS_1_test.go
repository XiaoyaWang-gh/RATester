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
			name: "test_0",
			args: args{
				text: "test_0",
			},
			want: template.HTML("<link href=\"test_0\" rel=\"stylesheet\" />"),
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
