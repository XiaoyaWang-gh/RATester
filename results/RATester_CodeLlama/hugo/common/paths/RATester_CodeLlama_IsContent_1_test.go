package paths

import (
	"fmt"
	"testing"
)

func TestIsContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{}
	p.bundleType = PathTypeContentResource
	if !p.IsContent() {
		t.Errorf("IsContent() = %v, want %v", false, true)
	}
}
