package hugolib

import (
	"fmt"
	"testing"
)

func TestisRenderedAny_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	po := &pageState{
		pageOutputs: []*pageOutput{
			{render: true},
			{render: false},
			{render: true},
		},
	}

	result := po.isRenderedAny()

	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}
