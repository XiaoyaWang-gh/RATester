package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestRender_1(t *testing.T) {
	type args struct {
		ctx converter.RenderContext
		doc any
	}

	tests := []struct {
		name    string
		c       *goldmarkConverter
		args    args
		want    converter.ResultRender
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

			got, err := tt.c.Render(tt.args.ctx, tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("goldmarkConverter.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goldmarkConverter.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
