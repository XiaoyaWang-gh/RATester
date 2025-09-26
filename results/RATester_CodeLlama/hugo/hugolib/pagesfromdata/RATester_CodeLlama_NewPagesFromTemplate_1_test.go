package pagesfromdata

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewPagesFromTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var opts PagesFromTemplateOptions
	var pt *PagesFromTemplate

	// Act
	pt = NewPagesFromTemplate(opts)

	// Assert
	assert.NotNil(t, pt)
}
