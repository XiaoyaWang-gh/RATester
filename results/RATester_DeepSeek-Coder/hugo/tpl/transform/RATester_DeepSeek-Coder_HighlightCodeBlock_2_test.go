package transform

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter/hooks"
	"github.com/gohugoio/hugo/markup/highlight"
)

func TestHighlightCodeBlock_2(t *testing.T) {
	type args struct {
		ctx  hooks.CodeblockContext
		opts []any
	}

	tests := []struct {
		name    string
		args    args
		want    highlight.HighlightResult
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

			ns := &Namespace{}
			got, err := ns.HighlightCodeBlock(tt.args.ctx, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.HighlightCodeBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.HighlightCodeBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
