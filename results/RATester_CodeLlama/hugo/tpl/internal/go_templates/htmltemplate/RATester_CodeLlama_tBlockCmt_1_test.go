package template

import (
	"fmt"
	"testing"
)

func TestTBlockCmt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := context{}
	s := []byte("")
	c, i := tBlockCmt(c, s)
	if i != len(s) {
		t.Errorf("tBlockCmt(%q, %q) = %d, want %d", c, s, i, len(s))
	}
}
