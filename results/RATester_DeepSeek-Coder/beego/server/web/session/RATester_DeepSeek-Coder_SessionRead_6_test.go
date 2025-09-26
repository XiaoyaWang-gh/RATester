package session

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestSessionRead_6(t *testing.T) {
	t.Parallel()
	fp := &FileProvider{
		lock:        sync.RWMutex{},
		maxlifetime: 100,
		savePath:    "/tmp/test",
	}
	ctx := context.Background()

	t.Run("valid sid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		sid := "valid_sid"
		_, err := fp.SessionRead(ctx, sid)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("invalid sid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		sid := "in/valid_sid"
		_, err := fp.SessionRead(ctx, sid)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("short sid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		sid := "sh"
		_, err := fp.SessionRead(ctx, sid)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
