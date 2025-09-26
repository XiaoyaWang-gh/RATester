package etag

import (
	"fmt"
	"testing"
)

func TestAppendUint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var dst []byte
	var n uint32
	dst = appendUint(dst, n)
	if len(dst) != 0 {
		t.Errorf("dst length is %d, want 0", len(dst))
	}
}
