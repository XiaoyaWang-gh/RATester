package tplimpl

import (
	"fmt"
	"testing"
)

func TestTemplateBaseName_1(t *testing.T) {
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
			name: "Test with shortcode type and name with internal path prefix",
			args: args{
				typ:  templateShortcode,
				name: internalPathPrefix + "test",
			},
			want: "test",
		},
		{
			name: "Test with shortcode type and name without internal path prefix",
			args: args{
				typ:  templateShortcode,
				name: "test",
			},
			want: "test",
		},
		{
			name: "Test with other type",
			args: args{
				typ:  2,
				name: "test",
			},
			want: "",
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
