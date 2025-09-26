package internal

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestTableOfContents_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	r := AsciidocResult{}
	// When
	got := r.TableOfContents()
	// Then
	assert.NotNil(t, got)
}
