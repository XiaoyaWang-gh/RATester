package gin

import (
	"fmt"
	"testing"
)

func TestWith_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{}

	testOpt := func(e *Engine) {
		e.RedirectTrailingSlash = true
	}

	engine.With(testOpt)

	if !engine.RedirectTrailingSlash {
		t.Error("Expected RedirectTrailingSlash to be true, but it was false")
	}
}
