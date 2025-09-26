package alils

import (
	"fmt"
	"testing"
)

func TestSize_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &LogGroupList{}
	n := m.Size()
	if n != 0 {
		t.Errorf("Size() = %v, want %v", n, 0)
	}
}
