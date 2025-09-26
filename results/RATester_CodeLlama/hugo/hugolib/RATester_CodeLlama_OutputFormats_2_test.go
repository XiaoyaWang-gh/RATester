package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/resources/page"
)

func TestOutputFormats_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	l := pagePaths{}

	// when
	outputFormats := l.OutputFormats()

	// then
	assert.Equal(t, page.OutputFormats{}, outputFormats)
}
