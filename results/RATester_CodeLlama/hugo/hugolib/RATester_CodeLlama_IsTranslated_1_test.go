package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestIsTranslated_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{}
	p.pageOutput = &pageOutput{}
	p.pageOutput.f = output.Format{
		Name: "en",
	}
	if !p.IsTranslated() {
		t.Errorf("IsTranslated() = false, want true")
	}
}
