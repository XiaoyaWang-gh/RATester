package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib/doctree"
)

func TestForEachResourceInPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var m pageMap
	var ps pageState
	var lockType doctree.LockType
	var exact bool
	var handle func(resourceKey string, n contentNodeI, match doctree.DimensionFlag) (bool, error)

	// Act
	err := m.forEachResourceInPage(&ps, lockType, exact, handle)

	// Assert
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
}
