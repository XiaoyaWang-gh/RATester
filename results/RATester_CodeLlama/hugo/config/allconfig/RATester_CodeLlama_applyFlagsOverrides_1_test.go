package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestApplyFlagsOverrides_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	cfg := config.New()
	cfg.Set("a", "a")
	cfg.Set("b", "b")
	cfg.Set("c", "c")

	l := configLoader{
		cfg: cfg,
	}

	l.applyFlagsOverrides(cfg)

	if cfg.Get("a") != "a" {
		t.Errorf("Expected a to be unchanged, but was %q", cfg.Get("a"))
	}

	if cfg.Get("b") != "b" {
		t.Errorf("Expected b to be unchanged, but was %q", cfg.Get("b"))
	}

	if cfg.Get("c") != "c" {
		t.Errorf("Expected c to be unchanged, but was %q", cfg.Get("c"))
	}
}
