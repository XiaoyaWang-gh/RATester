package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsRebuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var h *HugoSites

	assert.False(t, h.isRebuild())

	h.buildCounter.Store(1)
	assert.True(t, h.isRebuild())
}
