package template

import (
	"fmt"
	"reflect"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestComputeOutCtx_1(t *testing.T) {
	type args struct {
		c context
		t *template.Template
	}
	tests := []struct {
		name string
		args args
		want context
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

			e := &escaper{}
			if got := e.computeOutCtx(tt.args.c, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("escaper.computeOutCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
