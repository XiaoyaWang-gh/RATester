package pandoc

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestConvert_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &pandocConverter{
		ctx: converter.DocumentContext{
			DocumentID: "test",
		},
	}
	ctx := converter.RenderContext{
		Src: []byte("test"),
	}
	result, err := c.Convert(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result is nil")
	}
	if result.Bytes() == nil {
		t.Fatal("result.Bytes() is nil")
	}
}
