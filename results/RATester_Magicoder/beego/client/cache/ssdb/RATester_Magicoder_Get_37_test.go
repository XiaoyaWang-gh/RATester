package ssdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestGet_37(t *testing.T) {
	ctx := context.Background()
	rc := &Cache{
		conn:     &ssdb.Client{},
		conninfo: []string{},
	}

	t.Run("get value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "test_key"
		value := "test_value"
		rc.conn.Set(key, value)
		got, err := rc.Get(ctx, key)
		if err != nil {
			t.Errorf("Get() error = %v", err)
			return
		}
		if got != value {
			t.Errorf("Get() got = %v, want %v", got, value)
		}
	})

	t.Run("get non-existing value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "non_existing_key"
		got, err := rc.Get(ctx, key)
		if err == nil {
			t.Errorf("Get() error = %v, want non-nil", err)
			return
		}
		if got != nil {
			t.Errorf("Get() got = %v, want nil", got)
		}
	})
}
