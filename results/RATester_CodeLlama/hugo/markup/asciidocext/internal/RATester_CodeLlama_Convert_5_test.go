package internal

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestConvert_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &AsciidocConverter{}
	ctx := converter.RenderContext{
		Src: []byte(""),
	}
	_, err := a.Convert(ctx)
	if err != nil {
		t.Errorf("Convert() error = %v", err)
		return
	}
}
