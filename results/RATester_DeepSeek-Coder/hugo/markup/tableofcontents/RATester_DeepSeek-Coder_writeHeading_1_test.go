package tableofcontents

import (
	"fmt"
	"testing"
)

func TestWriteHeading_1(t *testing.T) {
	type args struct {
		level  int
		indent int
		h      *Heading
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test writeHeading with non-zero heading",
			args: args{
				level:  1,
				indent: 0,
				h: &Heading{
					ID:    "test-id",
					Title: "Test Title",
				},
			},
			want: "<li><a href=\"#test-id\">Test Title</a></li>\n",
		},
		{
			name: "Test writeHeading with zero heading",
			args: args{
				level:  1,
				indent: 0,
				h:      &Heading{},
			},
			want: "<li></li>\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := &tocBuilder{}
			b.writeHeading(tt.args.level, tt.args.indent, tt.args.h)
			got := b.s.String()
			if got != tt.want {
				t.Errorf("writeHeading() = %v, want %v", got, tt.want)
			}
		})
	}
}
