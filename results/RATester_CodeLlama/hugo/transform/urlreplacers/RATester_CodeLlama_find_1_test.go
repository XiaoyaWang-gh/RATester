package urlreplacers

import (
	"fmt"
	"testing"
)

func TestFind_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &prefix{}
	bs := []byte("")
	start := 0
	if got := p.find(bs, start); got != false {
		t.Errorf("find() = %v, want %v", got, false)
	}
}
