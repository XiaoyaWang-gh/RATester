package hugocontext

import (
	"fmt"
	"testing"
)

func TestKind_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoContext{}
	if h.Kind() != kindHugoContext {
		t.Errorf("expected %d, but got %d", kindHugoContext, h.Kind())
	}
}
