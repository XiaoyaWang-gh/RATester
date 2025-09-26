package mirror

import (
	"fmt"
	"testing"
)

func TestHijack_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := blackHoleResponseWriter{}
	_, _, err := b.Hijack()
	if err == nil {
		t.Error("expected error, got nil")
	}
}
