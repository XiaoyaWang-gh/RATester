package deps

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/tpl"
)

func TestStopStageRender_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	b := &BuildState{}
	stage := tpl.RenderingContext{}
	b.StartStageRender(stage)
	// When
	b.StopStageRender(stage)
	// Then
	assert.NotNil(t, b.DeferredExecutions)
	assert.NotNil(t, b.DeferredExecutionsGroupedByRenderingContext[stage])
	assert.Equal(t, b.DeferredExecutions, b.DeferredExecutionsGroupedByRenderingContext[stage])
}
