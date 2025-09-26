package strings

import (
	"fmt"
	"testing"
)

func TestCountWords_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("should return the number of words in a string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := "Hello, world!"
		expected := 2
		actual, err := ns.CountWords(input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})

	t.Run("should return the number of words in a string with CJK characters", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := "你好，世界！"
		expected := 2
		actual, err := ns.CountWords(input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})

	t.Run("should return an error if the input is not a string or convertible to string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := 123
		_, err := ns.CountWords(input)
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
}
