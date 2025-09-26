package tableofcontents

import (
	"fmt"
	"testing"
)

func TestWriteNav_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tocBuilder{}
	h := Headings{}
	b.writeNav(h)
	if b.s.String() != "<nav id=\"TableOfContents\"></nav>" {
		t.Errorf("writeNav() = %v, want %v", b.s.String(), "<nav id=\"TableOfContents\"></nav>")
	}
}
