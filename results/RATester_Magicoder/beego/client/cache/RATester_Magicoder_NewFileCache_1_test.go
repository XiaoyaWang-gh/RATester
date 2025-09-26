package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewFileCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := NewFileCache()

	if cache == nil {
		t.Error("Expected cache to be not nil")
	}

	ctx := context.Background()

	key := "testKey"
	value := "testValue"

	err := cache.Put(ctx, key, value, 1*time.Hour)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	got, err := cache.Get(ctx, key)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if got != value {
		t.Errorf("Expected %v, but got %v", value, got)
	}

	err = cache.Delete(ctx, key)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	got, err = cache.Get(ctx, key)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if got != nil {
		t.Errorf("Expected nil, but got %v", got)
	}
}
