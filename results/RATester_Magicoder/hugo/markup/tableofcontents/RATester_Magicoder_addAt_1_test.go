package tableofcontents

import (
	"fmt"
	"testing"
)

func TestaddAt_1(t *testing.T) {
	type args struct {
		h     *Heading
		row   int
		level int
	}
	tests := []struct {
		name string
		args args
		want *Fragments
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

			toc := &Fragments{}
			toc.addAt(tt.args.h, tt.args.row, tt.args.level)
		})
	}
}
