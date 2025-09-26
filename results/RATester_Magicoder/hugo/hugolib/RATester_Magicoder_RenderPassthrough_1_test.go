package hugolib

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/gohugoio/hugo/markup/converter/hooks"
)

func TestRenderPassthrough_1(t *testing.T) {
	type args struct {
		cctx context.Context
		w    io.Writer
		ctx  hooks.PassthroughContext
	}
	tests := []struct {
		name    string
		hr      hookRendererTemplate
		args    args
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

			if err := tt.hr.RenderPassthrough(tt.args.cctx, tt.args.w, tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("hookRendererTemplate.RenderPassthrough() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
