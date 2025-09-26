package page_generate

import (
	"fmt"
	"testing"
)

func TestImportsString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	imps := []string{"a", "b", "c"}
	want := "import (\n\"a\"\n\"b\"\n\"c\"\n)"
	got := importsString(imps)
	if got != want {
		t.Errorf("importsString(%v) = %q; want %q", imps, got, want)
	}
}
