package types

import (
	"fmt"
	"testing"
)

func TestAdd_15(t *testing.T) {
	t.Run("Adds a new value to the queue", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		q := &EvictingStringQueue{
			size: 2,
			vals: make([]string, 0),
			set:  make(map[string]bool),
		}

		q.Add("test1")

		if len(q.vals) != 1 {
			t.Errorf("Expected length of queue to be 1, got %d", len(q.vals))
		}

		if _, ok := q.set["test1"]; !ok {
			t.Errorf("Expected value 'test1' to be in set")
		}

		if q.vals[0] != "test1" {
			t.Errorf("Expected first value in queue to be 'test1', got %s", q.vals[0])
		}
	})

	t.Run("Evicts oldest value when queue is full", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		q := &EvictingStringQueue{
			size: 2,
			vals: []string{"test1", "test2"},
			set:  map[string]bool{"test1": true, "test2": true},
		}

		q.Add("test3")

		if len(q.vals) != 2 {
			t.Errorf("Expected length of queue to be 2, got %d", len(q.vals))
		}

		if _, ok := q.set["test1"]; ok {
			t.Errorf("Expected value 'test1' to be not in set")
		}

		if q.vals[0] != "test2" {
			t.Errorf("Expected first value in queue to be 'test2', got %s", q.vals[0])
		}

		if q.vals[1] != "test3" {
			t.Errorf("Expected second value in queue to be 'test3', got %s", q.vals[1])
		}
	})
}
