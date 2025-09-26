package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsApplicableTableForDB_1(t *testing.T) {
	t.Run("test case 1:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		val := reflect.ValueOf(1)
		db := "mysql"
		got := IsApplicableTableForDB(val, db)
		want := true
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("test case 2:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		val := reflect.ValueOf(1)
		db := "postgres"
		got := IsApplicableTableForDB(val, db)
		want := true
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("test case 3:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		val := reflect.ValueOf(1)
		db := "oracle"
		got := IsApplicableTableForDB(val, db)
		want := true
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("test case 4:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		val := reflect.ValueOf(1)
		db := "mssql"
		got := IsApplicableTableForDB(val, db)
		want := true
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("test case 5:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		val := reflect.ValueOf(1)
		db := "sqlite"
		got := IsApplicableTableForDB(val, db)
		want := true
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
