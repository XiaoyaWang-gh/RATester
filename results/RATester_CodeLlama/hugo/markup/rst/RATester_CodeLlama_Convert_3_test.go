package rst

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestConvert_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &rstConverter{}
	ctx := converter.RenderContext{}
	result, err := c.Convert(ctx)
	if err != nil {
		t.Errorf("Convert() error = %v", err)
		return
	}
	if result == nil {
		t.Errorf("Convert() result = nil")
		return
	}
	if result.Bytes() == nil {
		t.Errorf("Convert() result.Bytes() = nil")
		return
	}
}
