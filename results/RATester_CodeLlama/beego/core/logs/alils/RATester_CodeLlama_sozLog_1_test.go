package alils

import (
	"fmt"
	"testing"
)

func TestSozLog_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	x := uint64(1)
	n := sozLog(x)
	if n != 1 {
		t.Errorf("sozLog(%v) = %v, want %v", x, n, 1)
	}
}
