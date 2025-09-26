package config

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func TestGetNumWorkerMultiplier_1(t *testing.T) {
	t.Run("Test with environment variable set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("HUGO_NUMWORKERMULTIPLIER", "2")
		defer os.Unsetenv("HUGO_NUMWORKERMULTIPLIER")
		expected := 2
		actual := GetNumWorkerMultiplier()
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("Test with environment variable set to non-integer", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("HUGO_NUMWORKERMULTIPLIER", "non-integer")
		defer os.Unsetenv("HUGO_NUMWORKERMULTIPLIER")
		expected := runtime.NumCPU()
		actual := GetNumWorkerMultiplier()
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("Test with environment variable set to zero", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("HUGO_NUMWORKERMULTIPLIER", "0")
		defer os.Unsetenv("HUGO_NUMWORKERMULTIPLIER")
		expected := runtime.NumCPU()
		actual := GetNumWorkerMultiplier()
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("Test with environment variable not set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Unsetenv("HUGO_NUMWORKERMULTIPLIER")
		expected := runtime.NumCPU()
		actual := GetNumWorkerMultiplier()
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}
