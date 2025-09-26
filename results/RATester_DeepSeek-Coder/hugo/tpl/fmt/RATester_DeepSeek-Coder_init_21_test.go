package fmt

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestInit_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{}
	ctx := New(d)

	if ctx == nil {
		t.Errorf("Expected a non-nil context")
	}

	if ctx.Print("works") != "works" {
		t.Errorf("Expected 'works', got %v", ctx.Print("works"))
	}

	if ctx.Println("works") != "works\n" {
		t.Errorf("Expected 'works\\n', got %v", ctx.Println("works"))
	}

	if ctx.Printf("%s", "works") != "works" {
		t.Errorf("Expected 'works', got %v", ctx.Printf("%s", "works"))
	}

	if ctx.Errorf("%s", "failed") != "" {
		t.Errorf("Expected '', got %v", ctx.Errorf("%s", "failed"))
	}

	if ctx.Erroridf("my-err-id", "%s", "failed") != "" {
		t.Errorf("Expected '', got %v", ctx.Erroridf("my-err-id", "%s", "failed"))
	}

	if ctx.Warnidf("my-warn-id", "%s", "warning") != "" {
		t.Errorf("Expected '', got %v", ctx.Warnidf("my-warn-id", "%s", "warning"))
	}

	if ctx.Warnf("%s", "warning") != "" {
		t.Errorf("Expected '', got %v", ctx.Warnf("%s", "warning"))
	}
}
