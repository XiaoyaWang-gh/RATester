package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestVacuum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bc := &MemoryCache{
		dur: 10 * time.Second,
	}
	bc.vacuum()
}
