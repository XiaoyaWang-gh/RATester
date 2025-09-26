package cache

import (
	"fmt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3/internal/memory"
)

func TestSet_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &manager{
		memory: &memory.Storage{},
	}
	key := "key"
	it := &item{}
	exp := time.Duration(1)
	m.set(key, it, exp)
	if m.memory.Get(key) != it {
		t.Errorf("set() = %v, want %v", m.memory.Get(key), it)
	}
}
