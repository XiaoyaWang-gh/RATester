package paths

import (
	"fmt"
	"testing"
)

func TestIsLeafBundle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		bundleType: PathTypeLeaf,
	}
	if !p.IsLeafBundle() {
		t.Errorf("IsLeafBundle() = %v, want %v", p.IsLeafBundle(), true)
	}
}
