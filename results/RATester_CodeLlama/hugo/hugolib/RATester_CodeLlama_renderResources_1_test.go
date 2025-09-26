package hugolib

import (
	"fmt"
	"testing"
)

func TestRenderResources_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	p := &pageState{}

	// Act
	err := p.renderResources()

	// Assert
	if err != nil {
		t.Errorf("renderResources() error = %v", err)
		return
	}
}
