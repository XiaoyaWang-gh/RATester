package web

import (
	"fmt"
	"html/template"
	"testing"
)

func TestStr2html_1(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want template.HTML
	}{
		{
			name: "testStr2html",
			args: args{raw: "testStr2html"},
			want: template.HTML("testStr2html"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Str2html(tt.args.raw); got != tt.want {
				t.Errorf("Str2html() = %v, want %v", got, tt.want)
			}
		})
	}
}
