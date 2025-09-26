package tableofcontents

import (
	"fmt"
	"testing"
)

func TestWriteHeadings_1(t *testing.T) {
	type args struct {
		level   int
		indent  int
		h       Headings
		builder *tocBuilder
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := tt.args.builder
			b.writeHeadings(tt.args.level, tt.args.indent, tt.args.h)
			if got := b.s.String(); got != tt.want {
				t.Errorf("writeHeadings() = %v, want %v", got, tt.want)
			}
		})
	}
}
