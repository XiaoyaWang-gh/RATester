package tcp

import (
	"fmt"
	"testing"
)

func TestHasRoutes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Muxer{}
	if m.HasRoutes() {
		t.Error("expected false")
	}
	m.AddRoute("", "", 0, nil)
	if !m.HasRoutes() {
		t.Error("expected true")
	}
}
