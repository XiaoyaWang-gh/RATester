package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestWithRandomExpireOffsetFunc_1(t *testing.T) {
	t.Parallel()
	t.Run("test WithRandomExpireOffsetFunc", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		fn := func() time.Duration {
			return time.Duration(1)
		}
		opt := WithRandomExpireOffsetFunc(fn)
		cache := &RandomExpireCache{}
		opt(cache)
		if cache.offset == nil {
			t.Errorf("cache.offset is nil")
		}
	})
}
