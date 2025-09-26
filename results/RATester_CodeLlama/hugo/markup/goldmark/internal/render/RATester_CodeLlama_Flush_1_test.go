package render

import (
	"fmt"
	"testing"
)

func TestFlush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BufWriter{}
	if err := b.Flush(); err != nil {
		t.Errorf("Flush() = %v, want nil", err)
	}
}
