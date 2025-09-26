package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewPagePaginator_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	source := &pageState{}

	// Act
	actual := newPagePaginator(source)

	// Assert
	assert.NotNil(t, actual)
	assert.Equal(t, source, actual.source)
	assert.Equal(t, &pagePaginatorInit{}, actual.pagePaginatorInit)
}
