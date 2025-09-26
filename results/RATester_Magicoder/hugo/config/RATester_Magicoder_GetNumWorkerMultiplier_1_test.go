package config

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func TestGetNumWorkerMultiplier_1(t *testing.T) {
	t.Run("Test with valid environment variable", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("HUGO_NUMWORKERMULTIPLIER", "2")
		defer os.Unsetenv("HUGO_NUMWORKERMULTIPLIER")

		result := GetNumWorkerMultiplier()
		if result != 2 {
			t.Errorf("Expected 2, got %d", result)
		}
	})

	t.Run("Test with invalid environment variable", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("HUGO_NUMWORKERMULTIPLIER", "invalid")
		defer os.Unsetenv("HUGO_NUMWORKERMULTIPLIER")

		result := GetNumWorkerMultiplier()
		if result != runtime.NumCPU() {
			t.Errorf("Expected %d, got %d", runtime.NumCPU(), result)
		}
	})

	t.Run("Test with empty environment variable", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Unsetenv("HUGO_NUMWORKERMULTIPLIER")

		result := GetNumWorkerMultiplier()
		if result != runtime.NumCPU() {
			t.Errorf("Expected %d, got %d", runtime.NumCPU(), result)
		}
	})
}
