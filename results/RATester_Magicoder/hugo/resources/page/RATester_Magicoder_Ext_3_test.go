package page

import (
	"fmt"
	"testing"
)

func TestExt_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Ext() != "" {
		t.Errorf("Expected Ext() to return an empty string, but got %s", p.Ext())
	}
}
