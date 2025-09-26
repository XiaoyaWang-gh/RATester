package migration

import (
	"fmt"
	"testing"
)

func TestAddIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	index := &Index{Name: "test_index"}
	m.AddIndex(index)

	if len(m.Indexes) != 1 {
		t.Errorf("Expected 1 index, got %d", len(m.Indexes))
	}

	if m.Indexes[0].Name != "test_index" {
		t.Errorf("Expected index name to be 'test_index', got '%s'", m.Indexes[0].Name)
	}
}
