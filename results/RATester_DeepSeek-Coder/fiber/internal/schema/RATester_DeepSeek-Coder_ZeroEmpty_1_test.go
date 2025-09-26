package schema

import (
	"fmt"
	"testing"
)

func TestZeroEmpty_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Decoder{}
	d.ZeroEmpty(true)
	if d.zeroEmpty != true {
		t.Errorf("Expected ZeroEmpty to be true, got %v", d.zeroEmpty)
	}
}
