package session

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestSessionExist_8(t *testing.T) {
	ctx := context.Background()
	fp := &FileProvider{
		lock:        sync.RWMutex{},
		maxlifetime: 10,
		savePath:    "/tmp",
	}

	t.Run("session id length less than 2", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		exist, err := fp.SessionExist(ctx, "1")
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		if exist {
			t.Errorf("expected session not exist but got exist")
		}
	})

	t.Run("session id length equal to 2", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		exist, err := fp.SessionExist(ctx, "12")
		if err != nil {
			t.Errorf("expected nil error but got %v", err)
		}
		if exist {
			t.Errorf("expected session not exist but got exist")
		}
	})

	t.Run("session id length greater than 2", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		exist, err := fp.SessionExist(ctx, "123")
		if err != nil {
			t.Errorf("expected nil error but got %v", err)
		}
		if exist {
			t.Errorf("expected session not exist but got exist")
		}
	})
}
