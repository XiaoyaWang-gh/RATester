package models

import (
	"fmt"
	"testing"
)

func TestGet_6(t *testing.T) {
	mc := &ModelCache{
		cache: map[string]*ModelInfo{
			"table1": {
				Table: "table1",
			},
			"table2": {
				Table: "table2",
			},
		},
	}

	t.Run("should return model info and true if table exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		mi, ok := mc.Get("table1")
		if !ok {
			t.Fatal("expected ok to be true")
		}
		if mi.Table != "table1" {
			t.Fatalf("expected table to be 'table1', got '%s'", mi.Table)
		}
	})

	t.Run("should return nil and false if table does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		mi, ok := mc.Get("table3")
		if ok {
			t.Fatal("expected ok to be false")
		}
		if mi != nil {
			t.Fatalf("expected mi to be nil, got '%v'", mi)
		}
	})
}
