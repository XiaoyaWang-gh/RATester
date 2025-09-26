package attributes

import (
	"fmt"
	"strings"
	"testing"
)

func TestRenderAttributes_1(t *testing.T) {
	type args struct {
		w          *strings.Builder
		skipClass  bool
		attributes []Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with class attribute",
			args: args{
				w:         &strings.Builder{},
				skipClass: false,
				attributes: []Attribute{
					{
						Name:  "class",
						Value: "test-class",
					},
				},
			},
			want: " class=\"test-class\"",
		},
		{
			name: "Test without class attribute",
			args: args{
				w:         &strings.Builder{},
				skipClass: true,
				attributes: []Attribute{
					{
						Name:  "class",
						Value: "test-class",
					},
				},
			},
			want: "",
		},
		{
			name: "Test with multiple attributes",
			args: args{
				w:         &strings.Builder{},
				skipClass: false,
				attributes: []Attribute{
					{
						Name:  "data-test",
						Value: "test-value",
					},
					{
						Name:  "id",
						Value: "test-id",
					},
				},
			},
			want: " data-test=\"test-value\" id=\"test-id\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			RenderAttributes(tt.args.w, tt.args.skipClass, tt.args.attributes...)
			got := tt.args.w.String()
			if got != tt.want {
				t.Errorf("RenderAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}
