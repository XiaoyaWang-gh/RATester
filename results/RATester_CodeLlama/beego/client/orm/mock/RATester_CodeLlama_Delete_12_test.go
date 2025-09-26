package mock

import (
	"fmt"
	"testing"
)

func TestDelete_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	if _, err := d.Delete(); err != nil {
		t.Errorf("Delete() error = %v", err)
	}
}
