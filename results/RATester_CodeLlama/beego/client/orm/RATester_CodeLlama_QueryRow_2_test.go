package orm

import (
	"fmt"
	"testing"
)

func TestQueryRow_2(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		d := &stmtQueryLog{}
		args := []interface{}{}

		r := d.QueryRow(args...)

		if r == nil {
			t.Error("expected r to be not nil")
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		d := &stmtQueryLog{}
		args := []interface{}{}

		r := d.QueryRow(args...)

		if r != nil {
			t.Error("expected r to be nil")
		}
	})
}
