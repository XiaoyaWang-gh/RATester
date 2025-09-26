package logs

import (
	"fmt"
	"testing"
)

func TestColorByStatus_1(t *testing.T) {
	initColor()

	t.Run("Test 200 status code", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ColorByStatus(200)
		expected := "green"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test 300 status code", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ColorByStatus(300)
		expected := "white"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test 400 status code", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ColorByStatus(400)
		expected := "yellow"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test 500 status code", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ColorByStatus(500)
		expected := "red"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
