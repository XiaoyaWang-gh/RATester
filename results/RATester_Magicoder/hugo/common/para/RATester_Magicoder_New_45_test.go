package para

import (
	"fmt"
	"testing"
)

func TestNew_45(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	numWorkers := 5
	workers := New(numWorkers)

	if len(workers.sem) != numWorkers {
		t.Errorf("Expected semaphore to have capacity of %d, but got %d", numWorkers, len(workers.sem))
	}
}
