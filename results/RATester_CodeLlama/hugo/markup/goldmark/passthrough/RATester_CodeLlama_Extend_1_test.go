package passthrough

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
)

func TestExtend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var e *passthroughExtension
	var m goldmark.Markdown

	// act
	e.Extend(m)

	// assert
	// TODO: add assertions
}
