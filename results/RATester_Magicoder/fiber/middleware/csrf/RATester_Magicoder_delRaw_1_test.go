package csrf

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3/internal/memory"
)

func TestdelRaw_1(t *testing.T) {
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

	manager.setRaw(key, value, 0)

	manager.delRaw(key)

	_, ok := manager.memory.Get(key).([]byte)
	if ok {
		t.Error("Expected value to be deleted")
	}
}
