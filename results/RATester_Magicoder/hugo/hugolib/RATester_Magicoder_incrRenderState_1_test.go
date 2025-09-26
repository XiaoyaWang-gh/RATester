package hugolib

import (
	"fmt"
	"testing"
)

func TestincrRenderState_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	po := &pageOutput{
		renderState: 0,
		renderOnce:  false,
	}

	po.incrRenderState()

	if po.renderState != 1 {
		t.Errorf("Expected renderState to be 1, but got %d", po.renderState)
	}

	if !po.renderOnce {
		t.Errorf("Expected renderOnce to be true, but got false")
	}
}
