package gin

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSource_1(t *testing.T) {
	t.Run("Testing with valid inputs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		lines := [][]byte{[]byte("Hello"), []byte("World")}
		n := 1
		expected := []byte("World")
		result := source(lines, n)
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Testing with invalid n", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		lines := [][]byte{[]byte("Hello"), []byte("World")}
		n := 10
		expected := dunno
		result := source(lines, n)
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Testing with negative n", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		lines := [][]byte{[]byte("Hello"), []byte("World")}
		n := -1
		expected := dunno
		result := source(lines, n)
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
