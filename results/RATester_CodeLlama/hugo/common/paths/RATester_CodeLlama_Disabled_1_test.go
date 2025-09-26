package paths

import (
	"fmt"
	"testing"
)

func TestDisabled_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{}
	if p.Disabled() {
		t.Errorf("Disabled() = %v, want %v", p.Disabled(), false)
	}
}
