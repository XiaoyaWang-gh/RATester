package fake

import (
	"fmt"
	"testing"
)

func TestTracker_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Clientset{}
	if got := c.Tracker(); got != c.tracker {
		t.Errorf("Clientset.Tracker() = %v, want %v", got, c.tracker)
	}
}
