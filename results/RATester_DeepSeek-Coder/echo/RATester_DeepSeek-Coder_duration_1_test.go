package echo

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration_1(t *testing.T) {
	t.Run("Test duration with valid duration string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			ValueFunc: func(sourceParam string) string {
				return "1h30m"
			},
			ErrorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
				return fmt.Errorf("error: %v", message)
			},
		}
		var dest time.Duration
		b.duration("durationParam", &dest, true)
		expected := time.Hour + 30*time.Minute
		if dest != expected {
			t.Errorf("Expected %v, got %v", expected, dest)
		}
	})

	t.Run("Test duration with invalid duration string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			ValueFunc: func(sourceParam string) string {
				return "invalid"
			},
			ErrorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
				return fmt.Errorf("error: %v", message)
			},
		}
		var dest time.Duration
		b.duration("durationParam", &dest, true)
		if dest != 0 {
			t.Errorf("Expected 0, got %v", dest)
		}
	})

	t.Run("Test duration with empty string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			ValueFunc: func(sourceParam string) string {
				return ""
			},
			ErrorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
				return fmt.Errorf("error: %v", message)
			},
		}
		var dest time.Duration
		b.duration("durationParam", &dest, true)
		if dest != 0 {
			t.Errorf("Expected 0, got %v", dest)
		}
	})
}
