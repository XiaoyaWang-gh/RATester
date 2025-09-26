package goldmark

import (
	"fmt"
	"testing"
)

func TestDoc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := parserResult{
		doc: "hello",
	}
	if p.Doc() != "hello" {
		t.Errorf("expected %s, got %s", "hello", p.Doc())
	}
}
