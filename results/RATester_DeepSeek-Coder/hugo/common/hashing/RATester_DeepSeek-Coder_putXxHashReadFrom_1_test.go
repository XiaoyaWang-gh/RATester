package hashing

import (
	"fmt"
	"testing"

	"github.com/cespare/xxhash/v2"
)

func TestPutXxHashReadFrom_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &xxhashReadFrom{
		buff:   make([]byte, 1024),
		Digest: xxhash.New(),
	}

	putXxHashReadFrom(h)

	if h.buff != nil {
		t.Errorf("Expected buff to be nil, got %v", h.buff)
	}

	if h.Digest != nil {
		t.Errorf("Expected Digest to be nil, got %v", h.Digest)
	}
}
