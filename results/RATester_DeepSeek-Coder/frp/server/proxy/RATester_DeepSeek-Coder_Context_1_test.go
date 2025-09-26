package proxy

import (
	"context"
	"fmt"
	"testing"
)

func TestContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pxy := &BaseProxy{ctx: ctx}

	expected := ctx
	actual := pxy.Context()

	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
