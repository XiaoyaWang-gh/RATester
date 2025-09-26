package metric

import (
	"fmt"
	"testing"
)

func TestNewCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	counter := NewCounter()
	if counter.Count() != 0 {
		t.Errorf("counter.Count() = %d, want 0", counter.Count())
	}
	counter.Inc(1)
	if counter.Count() != 1 {
		t.Errorf("counter.Count() = %d, want 1", counter.Count())
	}
	counter.Inc(1)
	if counter.Count() != 2 {
		t.Errorf("counter.Count() = %d, want 2", counter.Count())
	}
	counter.Dec(1)
	if counter.Count() != 1 {
		t.Errorf("counter.Count() = %d, want 1", counter.Count())
	}
	counter.Dec(1)
	if counter.Count() != 0 {
		t.Errorf("counter.Count() = %d, want 0", counter.Count())
	}
	counter.Inc(1)
	if counter.Count() != 1 {
		t.Errorf("counter.Count() = %d, want 1", counter.Count())
	}
	counter.Inc(1)
	if counter.Count() != 2 {
		t.Errorf("counter.Count() = %d, want 2", counter.Count())
	}
	counter.Dec(1)
	if counter.Count() != 1 {
		t.Errorf("counter.Count() = %d, want 1", counter.Count())
	}
	counter.Dec(1)
	if counter.Count() != 0 {
		t.Errorf("counter.Count() = %d, want 0", counter.Count())
	}
}
