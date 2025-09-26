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
		t.Errorf("NewBuilder() = %v, want not nil", b)
	}

	if b.sourceRoots == nil {
		t.Errorf("NewBuilder().sourceRoots = %v, want not nil", b.sourceRoots)
	}
}
