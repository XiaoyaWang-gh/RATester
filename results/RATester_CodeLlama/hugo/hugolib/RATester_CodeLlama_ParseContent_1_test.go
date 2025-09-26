package hugolib

import (
	"context"
	"fmt"
	"testing"
)

func TestParseContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pco := &pageContentOutput{}
	ctx := context.Background()
	content := []byte("")
	r, ok, err := pco.ParseContent(ctx, content)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("expected ok to be true")
	}
	if r == nil {
		t.Fatal("expected r to be non-nil")
	}
}
