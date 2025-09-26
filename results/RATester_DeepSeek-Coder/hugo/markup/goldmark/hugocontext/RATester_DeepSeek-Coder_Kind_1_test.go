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

	ctx := &HugoContext{}
	expected := kindHugoContext
	actual := ctx.Kind()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
