package page

import (
	"fmt"
	"testing"
)

func TestIsHtmlIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagePathBuilder{}
	p.els = []string{"index.html"}
	if !p.IsHtmlIndex() {
		t.Errorf("IsHtmlIndex() = false, want true")
	}
}
