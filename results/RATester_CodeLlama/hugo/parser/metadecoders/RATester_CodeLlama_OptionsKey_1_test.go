package metadecoders

import (
	"fmt"
	"testing"
)

func TestOptionsKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := Decoder{
		Delimiter:  ',',
		Comment:    '#',
		LazyQuotes: true,
	}
	expected := "#,'true"
	if d.OptionsKey() != expected {
		t.Errorf("expected %s, got %s", expected, d.OptionsKey())
	}
}
