package tls

import (
	"fmt"
	"testing"

	"github.com/patrickmn/go-cache"
)

func TestResetCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := CertificateStore{
		CertCache: &cache.Cache{},
	}

	store.ResetCache()

	if store.CertCache != nil {
		t.Error("Expected cache to be flushed, but it wasn't")
	}
}
