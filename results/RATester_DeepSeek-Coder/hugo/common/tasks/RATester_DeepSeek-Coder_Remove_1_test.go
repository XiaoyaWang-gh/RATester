package tasks

import (
	"fmt"
	"testing"
)

func TestRemove_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{
		funcs: make(map[string]*Func),
	}

	r.funcs["testFunc"] = &Func{}

	r.Remove("testFunc")

	if _, ok := r.funcs["testFunc"]; ok {
		t.Errorf("Expected 'testFunc' to be removed from the funcs map")
	}
}
