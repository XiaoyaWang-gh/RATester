package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestDelete_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	fc := &FileCache{
		CachePath:      "./testdata",
		FileSuffix:     ".json",
		DirectoryLevel: 2,
		EmbedExpiry:    0,
	}
	// CONTEXT_1
	// CONTEXT_2
	ctx := context.Background()
	// CONTEXT_3
	// CONTEXT_4
	// CONTEXT_5
	key := "test_key"
	// CONTEXT_6
	// CONTEXT_7
	// CONTEXT_8
	// CONTEXT_9
	// TEST CASE
	err := fc.Put(ctx, key, "test_value", 0)
	if err != nil {
		t.Fatal(err)
	}
	// TEST CASE
	err = fc.Delete(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
	// TEST CASE
	_, err = fc.Get(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
}
