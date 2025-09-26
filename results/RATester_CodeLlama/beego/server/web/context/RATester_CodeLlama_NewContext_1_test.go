package context

import (
	"fmt"
	"testing"
)

func TestNewContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := NewContext()
	if ctx == nil {
		t.Error("NewContext() failed")
	}
}
