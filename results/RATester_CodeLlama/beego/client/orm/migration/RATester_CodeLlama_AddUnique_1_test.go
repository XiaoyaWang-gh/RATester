package migration

import (
	"fmt"
	"testing"
)

func TestAddUnique_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	unique := &Unique{}
	m.AddUnique(unique)
	if len(m.Uniques) != 1 {
		t.Errorf("Expected 1 unique, got %d", len(m.Uniques))
	}
}
