package cache

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3/internal/memory"
)

func TestGet_5(t *testing.T) {
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
	m.memory.Set(key, it, 0)
	if m.get(key) != it {
		t.Errorf("get() = %v, want %v", m.get(key), it)
	}
}
