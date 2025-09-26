package hugolib

import (
	"context"
	"fmt"
	"html/template"
	"testing"
)

func TestRenderString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pco := &pageContentOutput{
		po: &pageOutput{
			p: &pageState{
				pid: 1,
			},
		},
	}
	expected := template.HTML("<p>Hello, World!</p>")

	result, err := pco.RenderString(ctx)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
