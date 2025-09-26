package tplimpl

import (
	"fmt"
	"testing"
)

func TesttemplateBaseName_1(t *testing.T) {
	type args struct {
		typ  templateType
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				typ:  templateShortcode,
				name: "shortcodes/test.html",
			},
			want: "test.html",
		},
		{
			name: "Test case 2",
			args: args{
				typ:  templateShortcode,
				name: "shortcodes/test2/test.html",
			},
			want: "test2/test.html",
		},
		{
			name: "Test case 3",
			args: args{
				typ:  templateShortcode,
				name: "internal/shortcodes/test.html",
			},
			want: "test.html",
		},
		{
			name: "Test case 4",
			args: args{
				typ:  templateShortcode,
				name: "internal/shortcodes/test2/test.html",
			},
			want: "test2/test.html",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := templateBaseName(tt.args.typ, tt.args.name); got != tt.want {
				t.Errorf("templateBaseName() = %v, want %v", got, tt.want)
			}
		})
	}
}
