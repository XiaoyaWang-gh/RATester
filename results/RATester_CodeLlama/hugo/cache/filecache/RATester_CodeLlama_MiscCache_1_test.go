package filecache

import (
	"fmt"
	"testing"
)

func TestMiscCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := Caches{}
	if got := f.MiscCache(); got != nil {
		t.Errorf("MiscCache() = %v, want nil", got)
	}
}
