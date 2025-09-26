package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestQueryRow_6(t *testing.T) {

	t.Parallel()

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		query := "SELECT 1"
		args := []interface{}{}

		db := &DB{
			DB: &sql.DB{},
		}

		row := db.QueryRow(query, args...)

		if row.Err() != nil {
			t.Errorf("unexpected error: %v", row.Err())
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		query := "SELECT 1"
		args := []interface{}{}

		db := &DB{
			DB: &sql.DB{},
		}

		row := db.QueryRow(query, args...)

		if row.Err() == nil {
			t.Errorf("expected error")
		}
	})
}
