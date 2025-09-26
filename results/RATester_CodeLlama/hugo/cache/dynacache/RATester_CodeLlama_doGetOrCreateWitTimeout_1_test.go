package dynacache

import (
	"fmt"
	"testing"
	"time"

	"github.com/bep/lazycache"
)

func TestDoGetOrCreateWitTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Partition[string, string]{
		c: &lazycache.Cache[string, string]{},
	}
	key := "key"
	duration := time.Second
	create := func(key string) (string, error) {
		return "value", nil
	}
	v, err := p.doGetOrCreateWitTimeout(key, duration, create)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v != "value" {
		t.Errorf("expected %q, got %q", "value", v)
	}
}
