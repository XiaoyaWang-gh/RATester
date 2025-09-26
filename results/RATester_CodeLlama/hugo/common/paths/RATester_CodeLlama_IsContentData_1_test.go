package paths

import (
	"fmt"
	"testing"
)

func TestIsContentData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		bundleType: PathTypeContentData,
	}
	if !p.IsContentData() {
		t.Errorf("IsContentData() = %v, want %v", false, true)
	}
}
