package common

import (
	"fmt"
	"testing"
)

func TestGet_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pool := &copyBufferPool{}
	pool.New()
	buf := pool.Get()
	if len(buf) != PoolSizeCopy {
		t.Errorf("Get() = %v, want %v", len(buf), PoolSizeCopy)
	}
}
