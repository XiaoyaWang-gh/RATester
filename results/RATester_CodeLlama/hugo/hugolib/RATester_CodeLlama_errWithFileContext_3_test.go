package hugolib

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/source"
)

func TestErrWithFileContext_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	h := &HugoSites{}
	f := &source.File{}
	// When
	err := h.errWithFileContext(errors.New("test"), f)
	// Then
	assert.Equal(t, "test (in /)", err.Error())
}
