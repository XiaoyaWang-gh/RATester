package middleware

import (
	"fmt"
	"net/url"
	"testing"
)

func TestRemoveTarget_1(t *testing.T) {
	b := &commonBalancer{
		targets: []*ProxyTarget{
			{Name: "target1", URL: &url.URL{}},
			{Name: "target2", URL: &url.URL{}},
			{Name: "target3", URL: &url.URL{}},
		},
	}

	t.Run("remove existing target", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ok := b.RemoveTarget("target2")
		if !ok {
			t.Fatal("expected target to be removed")
		}
		if len(b.targets) != 2 {
			t.Fatalf("expected 2 targets, got %d", len(b.targets))
		}
		if b.targets[1].Name != "target3" {
			t.Fatalf("expected 'target3', got '%s'", b.targets[1].Name)
		}
	})

	t.Run("remove non-existing target", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ok := b.RemoveTarget("target4")
		if ok {
			t.Fatal("expected target not to be removed")
		}
		if len(b.targets) != 2 {
			t.Fatalf("expected 2 targets, got %d", len(b.targets))
		}
	})
}
