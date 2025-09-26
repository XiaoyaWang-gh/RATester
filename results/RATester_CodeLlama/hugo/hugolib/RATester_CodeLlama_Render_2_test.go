package hugolib

import (
	"context"
	"fmt"
	"testing"
)

func TestRender_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pco := &pageContentOutput{}
	layout := []string{"layout"}
	res, err := pco.Render(context.Background(), layout...)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	if res != "" {
		t.Errorf("expected empty result, got %q", res)
	}
}
