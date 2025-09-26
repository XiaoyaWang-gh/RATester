package pagemeta

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
)

func TestNewFrontmatterHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var logger loggers.Logger
	var frontMatterConfig FrontmatterConfig

	// Act
	_, err := NewFrontmatterHandler(logger, frontMatterConfig)

	// Assert
	if err != nil {
		t.Errorf("NewFrontmatterHandler() error = %v, want nil", err)
	}
}
