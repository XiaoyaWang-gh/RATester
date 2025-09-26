package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestSessionExist_4(t *testing.T) {
	ctx := context.Background()
	rp := &Provider{
		poollist: &redis.Client{},
	}

	t.Run("SessionExist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		rp.poollist.Set(ctx, "test_sid", "test_value", 0)
		exist, err := rp.SessionExist(ctx, "test_sid")
		if err != nil {
			t.Errorf("SessionExist() error = %v", err)
			return
		}
		if !exist {
			t.Errorf("SessionExist() = %v, want %v", exist, true)
		}
	})

	t.Run("SessionExist_NotExist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		exist, err := rp.SessionExist(ctx, "not_exist_sid")
		if err != nil {
			t.Errorf("SessionExist() error = %v", err)
			return
		}
		if exist {
			t.Errorf("SessionExist() = %v, want %v", exist, false)
		}
	})
}
