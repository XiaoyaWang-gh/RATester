package math

import (
	"fmt"
	"testing"
)

func TestCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	initialCounter := ns.Counter()
	expectedCounter := initialCounter + 1

	actualCounter := ns.Counter()
	if actualCounter != expectedCounter {
		t.Errorf("Expected counter to be %d, but got %d", expectedCounter, actualCounter)
	}
}
