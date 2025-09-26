package tcpmiddleware

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestNewBuilder_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	configs := map[string]*runtime.TCPMiddlewareInfo{
		"foo": {
			Status: "enabled",
			UsedBy: []string{"bar"},
		},
	}
	b := NewBuilder(configs)
	if b == nil {
		t.Fatal("expected a non-nil Builder")
	}
	if b.configs == nil {
		t.Fatal("expected a non-nil configs map")
	}
	if len(b.configs) != 1 {
		t.Fatalf("expected a configs map with 1 element, got %d", len(b.configs))
	}
	if b.configs["foo"] == nil {
		t.Fatal("expected a non-nil config for key 'foo'")
	}
	if b.configs["foo"].Status != "enabled" {
		t.Fatalf("expected config for key 'foo' to have status 'enabled', got %q", b.configs["foo"].Status)
	}
	if len(b.configs["foo"].UsedBy) != 1 {
		t.Fatalf("expected config for key 'foo' to have 1 usedBy element, got %d", len(b.configs["foo"].UsedBy))
	}
	if b.configs["foo"].UsedBy[0] != "bar" {
		t.Fatalf("expected config for key 'foo' to have usedBy element 'bar', got %q", b.configs["foo"].UsedBy[0])
	}
}
