package metrics

import (
	"fmt"
	"testing"
	"time"
)

func TestLen_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := bySum{
		{sum: 1 * time.Second},
		{sum: 2 * time.Second},
		{sum: 3 * time.Second},
	}

	expected := 3
	actual := b.Len()

	if actual != expected {
		t.Errorf("Expected Len() to return %d, but got %d", expected, actual)
	}
}
