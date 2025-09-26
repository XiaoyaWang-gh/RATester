package csrf

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3/internal/memory"
)

func TestgetRaw_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &storageManager{
		memory: &memory.Storage{},
	}

	key := "testKey"
	expectedValue := []byte("testValue")

	manager.memory.Set(key, expectedValue, 0)

	result := manager.getRaw(key)

	if !bytes.Equal(result, expectedValue) {
		t.Errorf("Expected %v, got %v", expectedValue, result)
	}
}
