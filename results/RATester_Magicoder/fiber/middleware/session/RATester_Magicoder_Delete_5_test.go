package session

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3/internal/storage/memory"
)

func TestDelete_5(t *testing.T) {
	store := &Store{
		Config: Config{
			Storage: memory.New(),
		},
	}

	t.Run("delete session with empty id", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := store.Delete("")
		if err != ErrEmptySessionID {
			t.Errorf("Expected error %v, got %v", ErrEmptySessionID, err)
		}
	})

	t.Run("delete session with non-empty id", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := store.Delete("non-empty-id")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
