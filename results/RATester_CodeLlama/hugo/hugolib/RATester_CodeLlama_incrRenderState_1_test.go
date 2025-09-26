package hugolib

import (
	"fmt"
	"testing"
)

func TestIncrRenderState_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	po := &pageOutput{}
	po.incrRenderState()
	if po.renderState != 1 {
		t.Errorf("renderState should be 1, but is %d", po.renderState)
	}
	if !po.renderOnce {
		t.Error("renderOnce should be true")
	}
}
