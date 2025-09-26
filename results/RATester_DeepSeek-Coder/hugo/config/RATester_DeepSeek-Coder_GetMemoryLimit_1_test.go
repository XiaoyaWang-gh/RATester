package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/pbnjay/memory"
)

func TestGetMemoryLimit_1(t *testing.T) {
	t.Run("Test with HUGO_MEMORYLIMIT set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("HUGO_MEMORYLIMIT", "1")
		defer os.Unsetenv("HUGO_MEMORYLIMIT")

		expected := uint64(1)
		result := GetMemoryLimit()

		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("Test with no HUGO_MEMORYLIMIT set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Unsetenv("HUGO_MEMORYLIMIT")

		expected := uint64(memory.TotalMemory() / 4)
		result := GetMemoryLimit()

		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
