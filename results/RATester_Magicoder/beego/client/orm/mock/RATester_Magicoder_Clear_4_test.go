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
	count, err := d.Clear()
	if err != nil {
		t.Errorf("Clear() error = %v, wantErr %v", err, nil)
		return
	}
	if count != 0 {
		t.Errorf("Clear() = %v, want %v", count, 0)
	}
}
