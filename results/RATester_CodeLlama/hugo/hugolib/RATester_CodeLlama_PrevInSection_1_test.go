package hugolib

import (
	"fmt"
	"testing"
)

func TestPrevInSection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := pagePositionInSection{}
	if got := p.PrevInSection(); got != nil {
		t.Errorf("PrevInSection() = %v, want nil", got)
	}
}
