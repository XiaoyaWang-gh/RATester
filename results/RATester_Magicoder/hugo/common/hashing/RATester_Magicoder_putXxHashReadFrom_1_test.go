package hashing

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cespare/xxhash/v2"
)

func TestputXxHashReadFrom_1(t *testing.T) {
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

	// Check if the object is back in the pool
	h2 := xXhashReadFromPool.Get().(*xxhashReadFrom)
	defer xXhashReadFromPool.Put(h2)

	if !reflect.DeepEqual(h, h2) {
		t.Errorf("Expected %v, got %v", h, h2)
	}
}
