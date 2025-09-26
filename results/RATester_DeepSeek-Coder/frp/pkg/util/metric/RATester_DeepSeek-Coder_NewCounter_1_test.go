package metric

import (
	"fmt"
	"testing"
)

func TestNewCounter_1(t *testing.T) {
	t.Run("NewCounter", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		counter := NewCounter()
		if counter == nil {
			t.Errorf("NewCounter() = %v, want not nil", counter)
		}
	})

	t.Run("Count", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		counter := NewCounter()
		if counter.Count() != 0 {
			t.Errorf("Count() = %v, want 0", counter.Count())
		}
	})

	t.Run("Inc", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		counter := NewCounter()
		counter.Inc(1)
		if counter.Count() != 1 {
			t.Errorf("Count() = %v, want 1", counter.Count())
		}
	})

	t.Run("Dec", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		counter := NewCounter()
		counter.Inc(2)
		counter.Dec(1)
		if counter.Count() != 1 {
			t.Errorf("Count() = %v, want 1", counter.Count())
		}
	})

	t.Run("Snapshot", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		counter := NewCounter()
		counter.Inc(1)
		snapshot := counter.Snapshot()
		if snapshot.Count() != 1 {
			t.Errorf("Snapshot.Count() = %v, want 1", snapshot.Count())
		}
	})

	t.Run("Clear", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		counter := NewCounter()
		counter.Inc(1)
		counter.Clear()
		if counter.Count() != 0 {
			t.Errorf("Count() = %v, want 0", counter.Count())
		}
	})
}
