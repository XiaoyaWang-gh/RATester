package cache

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3/internal/memory"
)

func TestgetRaw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &manager{
		memory: &memory.Storage{},
	}

	key := "testKey"
	value := []byte("testValue")
	manager.memory.Set(key, value, 0)

	raw := manager.getRaw(key)

	if !bytes.Equal(raw, value) {
		t.Errorf("Expected %v, got %v", value, raw)
	}
}
