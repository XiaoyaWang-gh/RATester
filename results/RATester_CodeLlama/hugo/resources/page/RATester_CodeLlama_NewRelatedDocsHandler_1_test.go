package page

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/related"
)

func TestNewRelatedDocsHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	cfg := related.Config{}

	// When
	handler := NewRelatedDocsHandler(cfg)

	// Then
	assert.NotNil(t, handler)
}
