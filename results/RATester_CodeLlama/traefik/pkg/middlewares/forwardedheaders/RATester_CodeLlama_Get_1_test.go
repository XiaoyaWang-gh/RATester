package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestGet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := unsafeHeader{}
	h.Set("key", "value")
	if h.Get("key") != "value" {
		t.Errorf("Get() = %q, want %q", h.Get("key"), "value")
	}
}
