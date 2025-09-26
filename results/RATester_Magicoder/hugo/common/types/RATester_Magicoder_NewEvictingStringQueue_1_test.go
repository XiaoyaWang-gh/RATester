package types

import (
	"fmt"
	"testing"
)

func TestNewEvictingStringQueue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	q := NewEvictingStringQueue(3)

	if q.Len() != 0 {
		t.Errorf("Expected queue length to be 0, got %d", q.Len())
	}

	q.Add("a")
	q.Add("b")
	q.Add("c")

	if q.Len() != 3 {
		t.Errorf("Expected queue length to be 3, got %d", q.Len())
	}

	if !q.Contains("a") || !q.Contains("b") || !q.Contains("c") {
		t.Errorf("Expected queue to contain 'a', 'b', 'c'")
	}

	if q.Peek() != "c" {
		t.Errorf("Expected peek to return 'c', got %s", q.Peek())
	}

	all := q.PeekAll()
	if len(all) != 3 || all[0] != "c" || all[1] != "b" || all[2] != "a" {
		t.Errorf("Expected PeekAll to return ['c', 'b', 'a'], got %v", all)
	}

	q.Add("d")

	if q.Len() != 3 {
		t.Errorf("Expected queue length to be 3 after adding 'd', got %d", q.Len())
	}

	if q.Contains("a") {
		t.Errorf("Expected queue not to contain 'a' after adding 'd'")
	}

	if q.Peek() != "d" {
		t.Errorf("Expected peek to return 'd' after adding 'd', got %s", q.Peek())
	}

	all = q.PeekAll()
	if len(all) != 3 || all[0] != "d" || all[1] != "c" || all[2] != "b" {
		t.Errorf("Expected PeekAll to return ['d', 'c', 'b'], got %v", all)
	}
}
