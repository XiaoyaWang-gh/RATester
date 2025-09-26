package pageparser

import (
	"fmt"
	"testing"
)

func TestCreateSectionHandlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	handlers := createSectionHandlers(l)
	if len(handlers.handlers) != 2 {
		t.Errorf("createSectionHandlers() = %v, want 2", len(handlers.handlers))
	}
}
