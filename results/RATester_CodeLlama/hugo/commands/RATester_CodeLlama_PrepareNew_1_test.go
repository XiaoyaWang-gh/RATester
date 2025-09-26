package commands

import (
	"fmt"
	"testing"
)

func TestPrepareNew_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &fileChangeDetector{}
	f.PrepareNew()
	if f.current == nil {
		t.Error("current map should not be nil")
	}
	if f.prev == nil {
		t.Error("prev map should not be nil")
	}
}
