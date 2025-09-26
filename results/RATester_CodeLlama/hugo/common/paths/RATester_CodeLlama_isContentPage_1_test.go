package paths

import (
	"fmt"
	"testing"
)

func TestIsContentPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{}
	p.bundleType = PathTypeContentSingle
	if !p.isContentPage() {
		t.Errorf("Expected true")
	}
}
