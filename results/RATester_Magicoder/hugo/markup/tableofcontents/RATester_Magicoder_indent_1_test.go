package tableofcontents

import (
	"fmt"
	"strings"
	"testing"
)

func Testindent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tocBuilder{
		s: strings.Builder{},
	}

	b.indent(3)

	expected := "    "
	actual := b.s.String()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
