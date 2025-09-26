package resources

import (
	"context"
	"fmt"
	"testing"
)

func TestTransformWithContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	e := &errorResource{}

	_, err := e.TransformWithContext(ctx)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
