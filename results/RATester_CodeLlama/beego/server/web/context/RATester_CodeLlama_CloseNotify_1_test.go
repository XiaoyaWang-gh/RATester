package context

import (
	"fmt"
	"testing"
)

func TestCloseNotify_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	if r.CloseNotify() != nil {
		t.Errorf("CloseNotify() = %v, want nil", r.CloseNotify())
	}
}
