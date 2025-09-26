package tableofcontents

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddAt_1(t *testing.T) {
	type args struct {
		h     *Heading
		row   int
		level int
	}
	tests := []struct {
		name string
		toc  *Fragments
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

			tt.toc.addAt(tt.args.h, tt.args.row, tt.args.level)
			if !reflect.DeepEqual(tt.toc, tt.want) {
				t.Errorf("Fragments.addAt() = %v, want %v", tt.toc, tt.want)
			}
		})
	}
}
