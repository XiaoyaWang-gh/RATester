package hugo

import (
	"fmt"
	"testing"
)

func TestString_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := VersionString("1.0.0")
	if h.String() != "1.0.0" {
		t.Errorf("expected %q, got %q", "1.0.0", h.String())
	}
}
