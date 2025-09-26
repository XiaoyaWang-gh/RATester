package binding

import (
	"fmt"
	"testing"
)

func TestLazyinit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &defaultValidator{}
	v.lazyinit()
	if v.validate == nil {
		t.Errorf("v.validate should not be nil")
	}
}
