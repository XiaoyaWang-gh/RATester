package paths

import (
	"fmt"
	"testing"
)

func TestIsBundle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		bundleType: PathTypeLeaf,
	}
	if !p.IsBundle() {
		t.Errorf("IsBundle() = %v, want %v", p.IsBundle(), true)
	}
}
