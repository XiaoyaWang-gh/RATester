package paths

import (
	"fmt"
	"testing"
)

func TestUnnormalized_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		unnormalized: &Path{},
	}
	if p.Unnormalized() != p.unnormalized {
		t.Errorf("Unnormalized() = %v, want %v", p.Unnormalized(), p.unnormalized)
	}
}
