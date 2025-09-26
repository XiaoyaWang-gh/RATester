package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestQuery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DB{
		DB: &sql.DB{},
	}
	query := "query"
	args := []interface{}{}
	rows, err := d.Query(query, args...)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	if rows.Err() != nil {
		t.Fatal(rows.Err())
	}
}
