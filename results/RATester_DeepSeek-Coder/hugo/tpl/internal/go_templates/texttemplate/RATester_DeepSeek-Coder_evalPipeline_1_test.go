package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEvalPipeline_1(t *testing.T) {
	type args struct {
		dot  reflect.Value
		pipe *parse.PipeNode
	}
	tests := []struct {
		name      string
		args      args
		wantValue reflect.Value
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
			if gotValue := s.evalPipeline(tt.args.dot, tt.args.pipe); !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("state.evalPipeline() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
