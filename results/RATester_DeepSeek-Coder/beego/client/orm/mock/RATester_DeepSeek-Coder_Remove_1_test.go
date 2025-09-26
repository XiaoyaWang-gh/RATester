package mock

import (
	"fmt"
	"testing"
)

func TestRemove_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQueryM2Mer{}
	_, err := d.Remove(1, 2, 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
