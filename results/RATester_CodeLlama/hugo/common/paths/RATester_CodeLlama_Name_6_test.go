package paths

import (
	"fmt"
	"testing"
)

func TestName_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c",
	}
	if p.Name() != "c" {
		t.Errorf("Expected %q but got %q", "c", p.Name())
	}
}
