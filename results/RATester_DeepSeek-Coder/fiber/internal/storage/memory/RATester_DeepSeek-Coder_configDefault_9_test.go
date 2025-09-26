package memory

import (
	"fmt"
	"testing"
	"time"
)

func TestConfigDefault_9(t *testing.T) {
	t.Parallel()

	ConfigDefault := Config{
		GCInterval: 10 * time.Second,
	}

	t.Run("no config provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cfg := configDefault()
		if cfg.GCInterval != ConfigDefault.GCInterval {
			t.Errorf("expected %v, got %v", ConfigDefault.GCInterval, cfg.GCInterval)
		}
	})

	t.Run("config provided with zero GCInterval", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cfg := configDefault(Config{GCInterval: 0})
		if cfg.GCInterval != ConfigDefault.GCInterval {
			t.Errorf("expected %v, got %v", ConfigDefault.GCInterval, cfg.GCInterval)
		}
	})

	t.Run("config provided with non-zero GCInterval", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := 20 * time.Second
		cfg := configDefault(Config{GCInterval: expected})
		if cfg.GCInterval != expected {
			t.Errorf("expected %v, got %v", expected, cfg.GCInterval)
		}
	})
}
