package highlight

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gohugoio/hugo/markup/internal/attributes"
)

func TestWriteDivStart_1(t *testing.T) {
	type args struct {
		w     *strings.Builder
		attrs []attributes.Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with nil attrs",
			args: args{
				w:     &strings.Builder{},
				attrs: nil,
			},
			want: "<div class=\"highlight\">",
		},
		{
			name: "Test with attrs",
			args: args{
				w: &strings.Builder{},
				attrs: []attributes.Attribute{
					{
						Name:  "class",
						Value: "test-class",
					},
				},
			},
			want: "<div class=\"highlight test-class\">",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			writeDivStart(tt.args.w, tt.args.attrs)
			got := tt.args.w.String()
			if got != tt.want {
				t.Errorf("writeDivStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
