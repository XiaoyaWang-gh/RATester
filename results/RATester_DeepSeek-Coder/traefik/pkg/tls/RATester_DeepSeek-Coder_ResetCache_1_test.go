package tls

import (
	"fmt"
	"testing"

	"github.com/patrickmn/go-cache"
)

func TestResetCache_1(t *testing.T) {
	testCases := []struct {
		desc     string
		cache    *cache.Cache
		expected *cache.Cache
	}{
		{
			desc:     "given a non-nil cache, when ResetCache is called, then the cache should be flushed",
			cache:    &cache.Cache{},
			expected: &cache.Cache{},
		},
		{
			desc:     "given a nil cache, when ResetCache is called, then the cache should remain nil",
			cache:    nil,
			expected: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			store := CertificateStore{
				CertCache: test.cache,
			}

			store.ResetCache()

			if test.cache != test.expected {
				t.Errorf("expected %v, got %v", test.expected, test.cache)
			}
		})
	}
}
