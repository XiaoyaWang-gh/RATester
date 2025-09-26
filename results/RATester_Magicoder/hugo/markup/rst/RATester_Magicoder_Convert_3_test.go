package rst

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestConvert_3(t *testing.T) {
	ctx := &rstConverter{
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
				Src: []byte("test"),
			},
			want:    converter.Bytes([]byte("test")),
			wantErr: false,
		},
		{
			name: "error",
			ctx: converter.RenderContext{
				Src: []byte("error"),
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
				t.Errorf("rstConverter.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rstConverter.Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
