package mock

import (
	"fmt"
	"testing"
)

func TestCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQueryM2Mer{}
	count, err := d.Count()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if count != 0 {
		t.Errorf("Expected count to be 0, but got %v", count)
	}
}
