package minifiers

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/media"
	"github.com/tdewolff/minify/v2"
)

func TestTransformer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var (
		m = Client{
			MinifyOutput: true,
			m:            minify.New(),
		}
		mediatype = media.Type{
			Type: "text/html",
		}
	)

	if got := m.Transformer(mediatype); got != nil {
		t.Errorf("Transformer() = %v, want nil", got)
	}
}
