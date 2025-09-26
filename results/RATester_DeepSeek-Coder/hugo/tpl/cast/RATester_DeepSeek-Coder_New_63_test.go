package cast

import (
	"fmt"
	"testing"
)

func TestNew_63(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := New()
	if ns == nil {
		t.Errorf("Expected a non-nil Namespace, got nil")
	}
}
