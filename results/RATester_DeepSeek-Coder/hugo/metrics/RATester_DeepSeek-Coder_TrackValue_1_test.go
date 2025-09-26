package metrics

import (
	"fmt"
	"testing"
)

func TestTrackValue_1(t *testing.T) {
	t.Parallel()

	store := &Store{
		calculateHints: true,
		diffs:          make(map[string]*diff),
		cached:         make(map[string]int),
	}

	t.Run("calculateHints is false", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		store.calculateHints = false
		store.TrackValue("key", "value", true)

		if _, found := store.diffs["key"]; found {
			t.Errorf("Expected 'key' to not be found in diffs")
		}
	})

	t.Run("calculateHints is true", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		store.calculateHints = true

		t.Run("cached is true", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			store.TrackValue("key", "value", true)

			if _, found := store.diffs["key"]; !found {
				t.Errorf("Expected 'key' to be found in diffs")
			}

			if store.cached["key"] != 1 {
				t.Errorf("Expected cached count for 'key' to be 1, got %d", store.cached["key"])
			}
		})

		t.Run("cached is false", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			store.TrackValue("key", "value", false)

			if _, found := store.diffs["key"]; !found {
				t.Errorf("Expected 'key' to be found in diffs")
			}

			if _, found := store.cached["key"]; found {
				t.Errorf("Expected 'key' to not be found in cached")
			}
		})
	})
}
