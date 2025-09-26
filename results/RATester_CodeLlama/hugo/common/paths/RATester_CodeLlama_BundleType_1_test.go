package paths

import (
	"fmt"
	"testing"
)

func TestBundleType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		bundleType: PathTypeBranch,
	}
	if p.BundleType() != PathTypeBranch {
		t.Errorf("expected %d, got %d", PathTypeBranch, p.BundleType())
	}
}
