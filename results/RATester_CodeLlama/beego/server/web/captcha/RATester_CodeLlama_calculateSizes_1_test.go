package captcha

import (
	"fmt"
	"testing"
)

func TestCalculateSizes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	m := &Image{}
	width := 100
	height := 100
	ncount := 10

	// Act
	m.calculateSizes(width, height, ncount)

	// Assert
	if m.dotSize != 1 {
		t.Errorf("dotSize should be 1, but was %d", m.dotSize)
	}
	if m.numWidth != 10 {
		t.Errorf("numWidth should be 10, but was %d", m.numWidth)
	}
	if m.numHeight != 10 {
		t.Errorf("numHeight should be 10, but was %d", m.numHeight)
	}
}
