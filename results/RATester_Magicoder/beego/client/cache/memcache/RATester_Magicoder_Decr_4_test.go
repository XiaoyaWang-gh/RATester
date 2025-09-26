package memcache

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
)

func TestDecr_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	mc := &Cache{
		conn: &memcache.Client{},
	}

	// Setup test data
	key := "testKey"
	value := int64(10)
	mc.conn.Set(&memcache.Item{Key: key, Value: []byte(strconv.FormatInt(value, 10))})

	// Test Decr
	err := mc.Decr(ctx, key)
	if err != nil {
		t.Errorf("Decr failed: %v", err)
	}

	// Verify result
	item, err := mc.conn.Get(key)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	result, err := strconv.ParseInt(string(item.Value), 10, 64)
	if err != nil {
		t.Errorf("ParseInt failed: %v", err)
	}
	if result != value-1 {
		t.Errorf("Expected %d, got %d", value-1, result)
	}
}
