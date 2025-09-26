package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := make(unsafeHeader)
	key := "key"
	value := "value"
	h.Set(key, value)
	if h[key][0] != value {
		t.Errorf("h.Set(%q, %q) = %q; want %q", key, value, h[key][0], value)
	}
}
