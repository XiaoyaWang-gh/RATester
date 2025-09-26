package web

import (
	"fmt"
	"testing"
)

func TestLen_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := ControllerCommentsSlice{}
	if p.Len() != 0 {
		t.Errorf("p.Len() = %v, want %v", p.Len(), 0)
	}
}
