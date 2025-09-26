package hugolib

import (
	"fmt"
	"testing"
)

func TestSetMetaPostPrepareRebuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	m := &pageMeta{}

	// Act
	m.setMetaPostPrepareRebuild()

	// Assert
	// ...
}
