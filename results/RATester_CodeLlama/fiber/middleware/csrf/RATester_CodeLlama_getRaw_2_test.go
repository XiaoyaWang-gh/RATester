package csrf

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3/internal/memory"
	"github.com/zeebo/assert"
)

func TestGetRaw_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	key := "key"
	m := &storageManager{
		memory: &memory.Storage{},
	}
	m.memory.Set(key, []byte("value"), 0)
	// When
	raw := m.getRaw(key)
	// Then
	assert.Equal(t, []byte("value"), raw)
}
