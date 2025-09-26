package paths

import (
	"fmt"
	"testing"
)

func TestIsBranchBundle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		bundleType: PathTypeBranch,
	}
	if !p.IsBranchBundle() {
		t.Errorf("IsBranchBundle() = %v, want %v", false, true)
	}
}
