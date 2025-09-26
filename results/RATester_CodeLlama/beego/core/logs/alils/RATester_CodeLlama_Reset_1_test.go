package alils

import (
	"fmt"
	"testing"
)

func TestReset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &LogContent{}
	m.Reset()
	if m.Key != nil {
		t.Errorf("m.Key should be nil")
	}
	if m.Value != nil {
		t.Errorf("m.Value should be nil")
	}
}
