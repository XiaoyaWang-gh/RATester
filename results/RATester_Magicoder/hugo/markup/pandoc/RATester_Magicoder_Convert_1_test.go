package pandoc

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestConvert_1(t *testing.T) {
	ctx := &pandocConverter{
		ctx: converter.DocumentContext{},
		cfg: converter.ProviderConfig{},
	}

	testCases := []struct {
		name    string
		ctx     converter.RenderContext
		want    converter.ResultRender
		wantErr bool
	}{
		{
			name: "success",
			ctx: converter.RenderContext{
				Src: []byte("content"),
			},
			want:    converter.Bytes([]byte("content")),
			wantErr: false,
		},
		{
			name: "error",
			ctx: converter.RenderContext{
				Src: []byte("content"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ctx.Convert(tt.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("pandocConverter.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pandocConverter.Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
