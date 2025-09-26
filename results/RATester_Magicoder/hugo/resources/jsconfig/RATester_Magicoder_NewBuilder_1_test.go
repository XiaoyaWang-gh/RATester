package jsconfig

import (
	"fmt"
	"testing"
)

func TestNewBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := NewBuilder()
	if b == nil {
		t.Error("NewBuilder() should not return nil")
	}
	if b.sourceRoots == nil {
		t.Error("NewBuilder() should initialize sourceRoots")
	}
}
