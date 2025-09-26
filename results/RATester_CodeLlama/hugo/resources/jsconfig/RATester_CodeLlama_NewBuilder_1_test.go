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

	builder := NewBuilder()
	if builder == nil {
		t.Error("NewBuilder() should not return nil")
	}
}
