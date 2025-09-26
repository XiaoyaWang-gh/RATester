package fiber

import (
	"fmt"
	"testing"
)

func TestReleaseRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	r := &Redirect{}

	// Act
	ReleaseRedirect(r)

	// Assert
	// ...
}
