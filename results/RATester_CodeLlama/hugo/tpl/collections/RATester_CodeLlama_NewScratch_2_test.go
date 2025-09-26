package collections

import (
	"fmt"
	"testing"
)

func TestNewScratch_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	if ns.NewScratch() == nil {
		t.Errorf("NewScratch() = nil, want not nil")
	}
}
