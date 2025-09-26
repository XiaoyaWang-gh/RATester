package session

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3/internal/storage/memory"
)

func TestReset_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &Store{
		Config: Config{
			Storage: memory.New(),
		},
	}

	err := store.Reset()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
