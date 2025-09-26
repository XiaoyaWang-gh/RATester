package goldmark

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func TestConvert_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &goldmarkConverter{
		md: goldmark.New(
			goldmark.WithExtensions(extension.GFM),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithXHTML(),
			),
		),
	}

	ctx := converter.RenderContext{
		Ctx: context.Background(),
		Src: []byte("Hello, **world**!"),
	}

	result, err := c.Convert(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(result.Bytes(), []byte("<p>Hello, <strong>world</strong>!</p>\n")) {
		t.Errorf("unexpected result: %q", result.Bytes())
	}
}
