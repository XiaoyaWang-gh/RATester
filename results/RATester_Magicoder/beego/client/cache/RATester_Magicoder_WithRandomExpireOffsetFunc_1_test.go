package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestWithRandomExpireOffsetFunc_1(t *testing.T) {
	testCases := []struct {
		name string
		fn   func() time.Duration
		want RandomExpireCacheOption
	}{
		{
			name: "Test case 1",
			fn: func() time.Duration {
				return time.Second * 10
			},
			want: func(cache *RandomExpireCache) {
				cache.offset = func() time.Duration {
					return time.Second * 10
				}
			},
		},
		{
			name: "Test case 2",
			fn: func() time.Duration {
				return time.Second * 20
			},
			want: func(cache *RandomExpireCache) {
				cache.offset = func() time.Duration {
					return time.Second * 20
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := WithRandomExpireOffsetFunc(tc.fn)
			cache := &RandomExpireCache{}
			got(cache)
			if cache.offset() != tc.fn() {
				t.Errorf("WithRandomExpireOffsetFunc() = %v, want %v", cache.offset(), tc.fn())
			}
		})
	}
}
