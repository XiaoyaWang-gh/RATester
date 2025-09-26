package related

import (
	"fmt"
	"testing"
)

func TestNewInvertedIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := Config{
		Threshold:    50,
		IncludeNewer: true,
		ToLower:      true,
		Indices: IndicesConfig{
			{Name: "index1"},
			{Name: "index2"},
		},
	}
	idx := NewInvertedIndex(cfg)

	if idx.cfg.Threshold != cfg.Threshold {
		t.Errorf("Expected Threshold to be %d, but got %d", cfg.Threshold, idx.cfg.Threshold)
	}
	if idx.cfg.IncludeNewer != cfg.IncludeNewer {
		t.Errorf("Expected IncludeNewer to be %t, but got %t", cfg.IncludeNewer, idx.cfg.IncludeNewer)
	}
	if idx.cfg.ToLower != cfg.ToLower {
		t.Errorf("Expected ToLower to be %t, but got %t", cfg.ToLower, idx.cfg.ToLower)
	}
	if len(idx.index) != len(cfg.Indices) {
		t.Errorf("Expected index length to be %d, but got %d", len(cfg.Indices), len(idx.index))
	}
	for _, conf := range cfg.Indices {
		if _, ok := idx.index[conf.Name]; !ok {
			t.Errorf("Expected index to contain %s, but it does not", conf.Name)
		}
	}
}
