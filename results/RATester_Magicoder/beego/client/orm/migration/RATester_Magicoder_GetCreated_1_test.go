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
		Created: "2022-01-01",
	}

	expected := int64(1640995200)
	actual := m.GetCreated()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
