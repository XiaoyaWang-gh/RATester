package context

import (
	"fmt"
	"testing"
)

func TestHijack_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	_, _, err := r.Hijack()
	if err == nil {
		t.Error("expected error")
	}
}
