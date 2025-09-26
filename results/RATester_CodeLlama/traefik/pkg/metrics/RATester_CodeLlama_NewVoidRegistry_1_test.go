package metrics

import (
	"fmt"
	"testing"
)

func TestNewVoidRegistry_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registry := NewVoidRegistry()
	if registry == nil {
		t.Error("NewVoidRegistry() should not return nil")
	}
}
