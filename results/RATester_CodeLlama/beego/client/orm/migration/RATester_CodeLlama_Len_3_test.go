package migration

import (
	"fmt"
	"testing"
)

func TestLen_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := dataSlice{
		data{
			created: 1,
			name:    "test",
			m:       nil,
		},
		data{
			created: 2,
			name:    "test",
			m:       nil,
		},
	}
	if d.Len() != 2 {
		t.Errorf("expected 2, got %d", d.Len())
	}
}
