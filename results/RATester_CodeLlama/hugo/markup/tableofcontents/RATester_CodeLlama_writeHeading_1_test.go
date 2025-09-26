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
		b    *tocBuilder
		args args
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

			tt.b.writeHeading(tt.args.level, tt.args.indent, tt.args.h)
		})
	}
}
