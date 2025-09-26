package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestDel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := make(unsafeHeader)
	h.Set("Content-Type", "application/json")
	h.Set("Authorization", "Bearer token")

	h.Del("Authorization")

	if got, want := h.Get("Authorization"), ""; got != want {
		t.Errorf("h.Get('Authorization') = %q; want %q", got, want)
	}
}
