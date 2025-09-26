package common

import (
	"fmt"
	"testing"
)

func TestGetBufPoolCopy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	BufPoolCopy.Put(make([]byte, PoolSizeCopy))
	buf := GetBufPoolCopy()
	if len(buf) != PoolSizeCopy {
		t.Errorf("Expected buffer size %d, got %d", PoolSizeCopy, len(buf))
	}
}
