package web

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDateParse_1(t *testing.T) {
	t.Run("Test Case 1:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		dateString := "2022-02-28 12:34:56"
		format := "2006-01-02 15:04:05"
		want := time.Date(2022, 2, 28, 12, 34, 56, 0, time.Local)
		got, err := DateParse(dateString, format)
		if err != nil {
			t.Errorf("DateParse() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DateParse() = %v, want %v", got, want)
		}
	})
	t.Run("Test Case 2:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		dateString := "2022-02-28 12:34:56"
		format := "2006-01-02 15:04:05"
		want := time.Date(2022, 2, 28, 12, 34, 56, 0, time.Local)
		got, err := DateParse(dateString, format)
		if err != nil {
			t.Errorf("DateParse() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DateParse() = %v, want %v", got, want)
		}
	})
}
