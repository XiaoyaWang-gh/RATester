package hugolib

import (
	"fmt"
	"testing"
)

func TestIsRenderedAny_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	po := &pageState{}
	po.pageOutputs = []*pageOutput{
		{
			render: true,
		},
		{
			render: false,
		},
		{
			render: true,
		},
	}

	// Act
	result := po.isRenderedAny()

	// Assert
	if result != true {
		t.Errorf("Expected isRenderedAny to return true, but it returned %v", result)
	}
}
