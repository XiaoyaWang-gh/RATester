package xlog

import (
	"fmt"
	"testing"
)

func TestNew_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := New()
	if l == nil {
		t.Error("New() return nil")
	}
}
