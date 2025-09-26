package hugo

import (
	"fmt"
	"html/template"
	"testing"
)

func TestGenerator_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := HugoInfo{
		CommitHash: "1234567890",
	}
	expected := template.HTML(`<meta name="generator" content="Hugo 0.0.0">`)
	actual := i.Generator()
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
