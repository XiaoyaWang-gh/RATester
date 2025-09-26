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

	size := 10
	queue := NewEvictingStringQueue(size)
	if queue.size != size {
		t.Errorf("Expected queue size to be %d, but got %d", size, queue.size)
	}
	if len(queue.set) != 0 {
		t.Errorf("Expected queue set to be empty, but got %v", queue.set)
	}
}
