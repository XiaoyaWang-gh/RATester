package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEnsurePipelineContains_1(t *testing.T) {
	type args struct {
		p *parse.PipeNode
		s []string
	}
	tests := []struct {
		name string
		args args
		want *parse.PipeNode
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

			ensurePipelineContains(tt.args.p, tt.args.s)
			if !reflect.DeepEqual(tt.args.p, tt.want) {
				t.Errorf("ensurePipelineContains() = %v, want %v", tt.args.p, tt.want)
			}
		})
	}
}
