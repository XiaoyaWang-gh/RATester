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
	i := []interface{}{}
	if _, err := d.Remove(i...); err != nil {
		t.Errorf("Remove() error = %v", err)
		return
	}
}
