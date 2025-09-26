package common

import (
	"fmt"
	"testing"
)

func TestNew_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var pool copyBufferPool
	pool.New()
	if pool.pool.New == nil {
		t.Error("pool.New is nil")
	}
}
