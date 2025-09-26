package web

import (
	"fmt"
	"testing"
)

func TestFinish_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Data: make(map[interface{}]interface{}),
	}
	ctrl.Finish()
	if len(ctrl.Data) != 0 {
		t.Errorf("Expected empty map, got %v", ctrl.Data)
	}
}
