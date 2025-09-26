package encoding

import (
	"fmt"
	"testing"
)

func TestNew_48(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := New()
	if ns == nil {
		t.Error("Expected a non-nil Namespace, got nil")
	}
}
