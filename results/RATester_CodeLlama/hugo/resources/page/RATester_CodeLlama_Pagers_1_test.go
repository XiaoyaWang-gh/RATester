package page

import (
	"fmt"
	"testing"
)

func TestPagers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{}
	if got := p.Pagers(); got != nil {
		t.Errorf("Paginator.Pagers() = %v, want nil", got)
	}
}
