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
	count, err := d.Delete()
	if err != nil {
		t.Errorf("Delete() error = %v, wantErr %v", err, nil)
		return
	}
	if count != 0 {
		t.Errorf("Delete() count = %v, want %v", count, 0)
	}
}
