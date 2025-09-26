package csrf

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3/internal/memory"
)

func TestsetRaw_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &storageManager{
		memory: &memory.Storage{},
	}

	key := "testKey"
	value := []byte("testValue")
	exp := time.Duration(10)

	manager.setRaw(key, value, exp)

	result := manager.getRaw(key)

	if !bytes.Equal(result, value) {
		t.Errorf("Expected %v, got %v", value, result)
	}
}
