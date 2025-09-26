package paths

import (
	"fmt"
	"testing"
)

func TestModifyPathBundleTypeResource_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{}
	ModifyPathBundleTypeResource(p)
	if p.bundleType != PathTypeContentResource {
		t.Errorf("p.bundleType = %v, want %v", p.bundleType, PathTypeContentResource)
	}
}
