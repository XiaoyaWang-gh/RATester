package cast

import (
	"fmt"
	"html/template"
	"reflect"
	"testing"
)

func TestConvertTemplateToString_1(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "case 1",
			args: args{
				v: template.HTML("html"),
			},
			want: "html",
		},
		{
			name: "case 2",
			args: args{
				v: template.CSS("css"),
			},
			want: "css",
		},
		{
			name: "case 3",
			args: args{
				v: template.HTMLAttr("htmlattr"),
			},
			want: "htmlattr",
		},
		{
			name: "case 4",
			args: args{
				v: template.JS("js"),
			},
			want: "js",
		},
		{
			name: "case 5",
			args: args{
				v: template.JSStr("jssstr"),
			},
			want: "jssstr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertTemplateToString(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTemplateToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
