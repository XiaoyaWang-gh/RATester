package cache

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"
)

func TestGet_18(t *testing.T) {
	fc := &FileCache{
		CachePath: "/tmp/test",
	}

	ctx := context.Background()

	t.Run("key not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := fc.Get(ctx, "not_exist")
		if err != ErrKeyExpired {
			t.Errorf("expected ErrKeyExpired, got %v", err)
		}
	})

	t.Run("key expired", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "expired"
		item := &FileCacheItem{
			Data:       "data",
			Lastaccess: time.Now(),
			Expired:    time.Now().Add(-time.Hour),
		}
		data, _ := GobEncode(item)
		ioutil.WriteFile(filepath.Join(fc.CachePath, key), data, 0644)

		_, err := fc.Get(ctx, key)
		if err != ErrKeyExpired {
			t.Errorf("expected ErrKeyExpired, got %v", err)
		}
	})

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "success"
		item := &FileCacheItem{
			Data:       "data",
			Lastaccess: time.Now(),
			Expired:    time.Now().Add(time.Hour),
		}
		data, _ := GobEncode(item)
		ioutil.WriteFile(filepath.Join(fc.CachePath, key), data, 0644)

		val, err := fc.Get(ctx, key)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if val != "data" {
			t.Errorf("expected 'data', got %v", val)
		}
	})
}
