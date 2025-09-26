package pagemeta

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewDefaultFrontmatterConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	expected := FrontmatterConfig{
		Date:        []string{fmDate, fmPubDate, fmLastmod},
		Lastmod:     []string{fmGitAuthorDate, fmLastmod, fmDate, fmPubDate},
		PublishDate: []string{fmPubDate, fmDate},
		ExpiryDate:  []string{fmExpiryDate},
	}

	// Act
	actual := newDefaultFrontmatterConfig()

	// Assert
	assert.Equal(t, expected, actual)
}
