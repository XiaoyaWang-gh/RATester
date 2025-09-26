package tplimpl

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	texttemplate "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestTrackDependencies_1(t *testing.T) {
	type args struct {
		ctx      context.Context
		tmpl     texttemplate.Preparer
		name     string
		receiver reflect.Value
	}
	tests := []struct {
		name string
		t    *templateExecHelper
		args args
		want context.Context
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

			if got := tt.t.trackDependencies(tt.args.ctx, tt.args.tmpl, tt.args.name, tt.args.receiver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("templateExecHelper.trackDependencies() = %v, want %v", got, tt.want)
			}
		})
	}
}
