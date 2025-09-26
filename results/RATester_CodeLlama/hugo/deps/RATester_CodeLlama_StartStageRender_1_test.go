package deps

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl"
)

func TestStartStageRender_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	var b *BuildState
	var stage tpl.RenderingContext

	// When
	b.StartStageRender(stage)

	// Then
	// TODO: add assertions
}
