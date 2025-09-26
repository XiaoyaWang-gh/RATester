package memcache

import (
	"context"
	"fmt"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
)

func TestDecr_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx := context.Background()
	mc := memcache.New("localhost:11211")
	rc := &Cache{conn: mc, conninfo: []string{"localhost:11211"}}

	key := "test_key"
	value := 10
	err := rc.Put(ctx, key, value, 0)
	if err != nil {
		t.Fatalf("Failed to put value: %v", err)
	}

	err = rc.Decr(ctx, key)
	if err != nil {
		t.Fatalf("Failed to decrease value: %v", err)
	}

	item, err := mc.Get(key)
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}

	var newValue int
	_, err = fmt.Sscanf(string(item.Value), "%d", &newValue)
	if err != nil {
		t.Fatalf("Failed to scan value: %v", err)
	}

	if newValue != value-1 {
		t.Errorf("Expected value to be %d, got %d", value-1, newValue)
	}
}
