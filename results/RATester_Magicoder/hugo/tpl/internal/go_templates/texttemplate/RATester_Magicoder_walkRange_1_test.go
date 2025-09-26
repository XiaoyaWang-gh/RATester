package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestwalkRange_1(t *testing.T) {
	tests := []struct {
		name string
		node *parse.RangeNode
		want any
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

			s := &state{}
			s.walkRange(reflect.ValueOf(tt.node), tt.node)
			// TODO: Add assertions.
		})
	}
}
