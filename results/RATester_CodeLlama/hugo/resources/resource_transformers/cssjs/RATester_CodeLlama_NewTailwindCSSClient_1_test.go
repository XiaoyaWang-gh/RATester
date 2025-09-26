package cssjs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/resources"
)

func TestNewTailwindCSSClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var rs *resources.Spec

	// Act
	var tailwindCSSClient *TailwindCSSClient
	tailwindCSSClient = NewTailwindCSSClient(rs)

	// Assert
	assert.NotNil(t, tailwindCSSClient)
}
