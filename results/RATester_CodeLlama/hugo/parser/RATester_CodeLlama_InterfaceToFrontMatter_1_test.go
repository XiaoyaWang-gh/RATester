package parser

import (
	"fmt"
	"io"
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
)

func TestInterfaceToFrontMatter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var in any
	var format metadecoders.Format
	var w io.Writer

	err := InterfaceToFrontMatter(in, format, w)
	if err == nil {
		t.Errorf("InterfaceToFrontMatter() expected error, got nil")
	}
}
