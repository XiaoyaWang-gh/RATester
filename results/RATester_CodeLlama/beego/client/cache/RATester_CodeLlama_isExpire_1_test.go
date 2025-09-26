package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestIsExpire_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mi := &MemoryItem{
		val:         "test",
		createdTime: time.Now(),
		lifespan:    10 * time.Second,
	}
	if mi.isExpire() {
		t.Error("MemoryItem is not expire")
	}
	time.Sleep(11 * time.Second)
	if !mi.isExpire() {
		t.Error("MemoryItem is expire")
	}
}
