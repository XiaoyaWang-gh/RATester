package migration

import (
	"fmt"
	"testing"
)

func TestGetCreated_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{
		Created: "2019-01-01 00:00:00",
	}
	if m.GetCreated() != 1546300800 {
		t.Errorf("GetCreated() = %v, want %v", m.GetCreated(), 1546300800)
	}
}
