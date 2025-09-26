package pagesfromdata

import (
	"fmt"
	"testing"
)

func TestSite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagesFromDataTemplateContext{}
	if p.Site() != nil {
		t.Errorf("expected nil")
	}
}
