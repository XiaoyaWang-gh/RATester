package orm

import (
	"fmt"
	"testing"
)

func TestGetDB_1(t *testing.T) {
	t.Run("TestGetDBWithAlias", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		db, err := GetDB("test")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if db == nil {
			t.Errorf("Expected db not to be nil")
		}
	})

	t.Run("TestGetDBWithoutAlias", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		db, err := GetDB()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if db == nil {
			t.Errorf("Expected db not to be nil")
		}
	})
}
