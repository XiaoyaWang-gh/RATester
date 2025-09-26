package page

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/media"
)

func TestExtractSummaryFromHTMLWithDivider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var mt media.Type
	var input string
	var divider string
	var result HtmlSummary

	// Act
	result = ExtractSummaryFromHTMLWithDivider(mt, input, divider)

	// Assert
	assert.Equal(t, result.source, input)
	assert.Equal(t, result.Divider.Low, strings.Index(input, divider))
	assert.Equal(t, result.Divider.High, result.Divider.Low+len(divider))

	if result.Divider.Low == -1 {
		return
	}

	ptag := result.resolveParagraphTagAndSetWrapper(mt)

	if !mt.IsHTML() {
		result.Divider, result.SummaryEndTag = expandSummaryDivider(result.source, ptag, result.Divider)
	}

	assert.Equal(t, result.SummaryLowHigh.Low, result.WrapperStart.High)
	assert.Equal(t, result.SummaryLowHigh.High, result.Divider.Low)
}
