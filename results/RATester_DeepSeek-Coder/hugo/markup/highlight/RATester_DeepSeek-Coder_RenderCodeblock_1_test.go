package highlight

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
	"github.com/gohugoio/hugo/markup/converter/hooks"
)

func TestRenderCodeblock_1(t *testing.T) {
	type args struct {
		cctx context.Context
		w    hugio.FlexiWriter
		ctx  hooks.CodeblockContext
	}
	tests := []struct {
		name    string
		h       chromaHighlighter
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

			err := tt.h.RenderCodeblock(tt.args.cctx, tt.args.w, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderCodeblock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
