package identity

import (
	"fmt"
	"testing"
)

func TestCheckMaxDepth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &Finder{}
	sid := &searchID{}
	level := 1
	if got := f.checkMaxDepth(sid, level); got != -1 {
		t.Errorf("checkMaxDepth() = %v, want -1", got)
	}
}
