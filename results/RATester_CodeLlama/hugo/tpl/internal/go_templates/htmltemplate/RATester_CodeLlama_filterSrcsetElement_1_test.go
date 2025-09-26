package template

import (
	"fmt"
	"strings"
	"testing"
)

func TestFilterSrcsetElement_1(t *testing.T) {
	type args struct {
		s     string
		left  int
		right int
		b     *strings.Builder
	}
	tests := []struct {
		name string
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

			filterSrcsetElement(tt.args.s, tt.args.left, tt.args.right, tt.args.b)
		})
	}
}
