package mock

import (
	"fmt"
	"testing"
)

func TestClear_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQueryM2Mer{}
	if _, err := d.Clear(); err != nil {
		t.Errorf("Clear() error = %v", err)
	}
}
