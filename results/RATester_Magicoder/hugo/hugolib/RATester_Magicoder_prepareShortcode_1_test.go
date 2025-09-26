package hugolib

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl"
)

func TestprepareShortcode_1(t *testing.T) {
	type args struct {
		ctx            context.Context
		level          int
		s              *Site
		tplVariants    tpl.TemplateVariants
		sc             *shortcode
		parent         *ShortcodeWithPage
		p              *pageState
		isRenderString bool
	}
	tests := []struct {
		name    string
		args    args
		want    shortcodeRenderer
		wantErr bool
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

			got, err := prepareShortcode(tt.args.ctx, tt.args.level, tt.args.s, tt.args.tplVariants, tt.args.sc, tt.args.parent, tt.args.p, tt.args.isRenderString)
			if (err != nil) != tt.wantErr {
				t.Errorf("prepareShortcode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareShortcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
