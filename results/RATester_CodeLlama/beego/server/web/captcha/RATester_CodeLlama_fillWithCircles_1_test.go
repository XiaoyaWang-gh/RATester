package captcha

import (
	"fmt"
	"testing"
)

func TestFillWithCircles_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	m := &Image{}
	n := 10
	maxradius := 10

	// Act
	m.fillWithCircles(n, maxradius)

	// Assert
	// TODO: Add assertions
}
