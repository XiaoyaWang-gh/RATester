package highlight

import (
	"fmt"
	"testing"
)

func TestApplyOptionsFromMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	optsm := map[string]any{
		"style": "github",
	}
	cfg := &Config{}
	err := applyOptionsFromMap(optsm, cfg)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Style != "github" {
		t.Errorf("cfg.Style = %q, want %q", cfg.Style, "github")
	}
}
