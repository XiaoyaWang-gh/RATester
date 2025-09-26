package tasks

import (
	"fmt"
	"testing"
)

func TestHas_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{
		funcs: make(map[string]*Func),
	}

	r.funcs["testFunc"] = &Func{}

	if !r.Has("testFunc") {
		t.Error("Expected to find 'testFunc'")
	}

	if r.Has("nonExistentFunc") {
		t.Error("Expected not to find 'nonExistentFunc'")
	}
}
