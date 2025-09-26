package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestNewParserContext_1(t *testing.T) {
	type args struct {
		rctx converter.RenderContext
	}
	tests := []struct {
		name string
		args args
		want *parserContext
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

			c := &goldmarkConverter{}
			if got := c.newParserContext(tt.args.rctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goldmarkConverter.newParserContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
