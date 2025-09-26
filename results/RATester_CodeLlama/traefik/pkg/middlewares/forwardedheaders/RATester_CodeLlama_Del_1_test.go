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

	h := unsafeHeader{}
	h.Set("key", "value")
	h.Del("key")
	if h.Get("key") != "" {
		t.Errorf("Del failed")
	}
}
