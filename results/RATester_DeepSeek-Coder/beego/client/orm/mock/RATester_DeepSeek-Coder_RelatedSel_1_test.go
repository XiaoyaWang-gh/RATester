package mock

import (
	"fmt"
	"testing"
)

func TestRelatedSel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	qs := d.RelatedSel()
	if qs != d {
		t.Errorf("Expected query seter to be the same as the input")
	}
}
