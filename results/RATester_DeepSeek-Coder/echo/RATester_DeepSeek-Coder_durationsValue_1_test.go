package echo

import (
	"fmt"
	"testing"
	"time"
)

func TestDurationsValue_1(t *testing.T) {
	t.Run("Test durationsValue method", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			ValuesFunc: func(sourceParam string) []string {
				if sourceParam == "duration" {
					return []string{"1s", "2s", "3s"}
				}
				return []string{}
			},
			ErrorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
				return fmt.Errorf("error: %v", message)
			},
		}

		t.Run("Test when values exist", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			dest := []time.Duration{}
			b.durationsValue("duration", &dest, true)
			if len(dest) != 3 {
				t.Errorf("Expected 3 durations, got %d", len(dest))
			}
		})

		t.Run("Test when values do not exist", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			dest := []time.Duration{}
			b.durationsValue("no_duration", &dest, true)
			if len(dest) != 0 {
				t.Errorf("Expected 0 durations, got %d", len(dest))
			}
			if b.errors == nil {
				t.Errorf("Expected error, got nil")
			}
		})
	})
}
