package process

import (
	"fmt"
	"testing"
)

func TestSetBeforeStopHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Process{}
	p.SetBeforeStopHandler(func() {
		fmt.Println("before stop handler")
	})
	if p.beforeStopHandler == nil {
		t.Error("beforeStopHandler is nil")
	}
}
