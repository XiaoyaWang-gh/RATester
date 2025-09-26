package redis

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestSessionRead_4(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	poollist := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rp := &Provider{
		maxlifetime: 1800,
		poollist:    poollist,
	}

	t.Run("SessionRead_Success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		sid := "test_session_id"
		expectedSessionStore := &SessionStore{
			p:           poollist,
			sid:         sid,
			values:      make(map[interface{}]interface{}),
			maxlifetime: 1800,
		}

		sessionStore, err := rp.SessionRead(ctx, sid)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !reflect.DeepEqual(sessionStore, expectedSessionStore) {
			t.Errorf("Expected session store %v, got %v", expectedSessionStore, sessionStore)
		}
	})

	t.Run("SessionRead_RedisError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		sid := "non_existent_session_id"
		_, err := rp.SessionRead(ctx, sid)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("SessionRead_DecodeError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		sid := "invalid_session_id"
		_, err := rp.SessionRead(ctx, sid)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
