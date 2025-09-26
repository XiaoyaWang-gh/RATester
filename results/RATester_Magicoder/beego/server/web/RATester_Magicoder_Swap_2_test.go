package web

import (
	"fmt"
	"testing"
)

func TestSwap_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice := ControllerCommentsSlice{
		ControllerComments{Method: "Method1"},
		ControllerComments{Method: "Method2"},
	}
	slice.Swap(0, 1)
	if slice[0].Method != "Method2" || slice[1].Method != "Method1" {
		t.Errorf("Swap failed. Expected: Method2, Method1. Got: %s, %s", slice[0].Method, slice[1].Method)
	}
}
