package fiber

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConvert_1(t *testing.T) {
	t.Run("successful conversion", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := "42"
		convertor := func(s string) (int, error) {
			return strconv.Atoi(s)
		}
		expected := 42
		result, err := Convert(value, convertor)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("failed conversion", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := "not an int"
		convertor := func(s string) (int, error) {
			return strconv.Atoi(s)
		}
		expectedErr := "failed to convert: strconv.Atoi: parsing \"not an int\": invalid syntax"
		_, err := Convert(value, convertor)
		if err == nil {
			t.Errorf("expected error: %v", expectedErr)
		} else if err.Error() != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("successful conversion with default value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := "not an int"
		convertor := func(s string) (int, error) {
			return strconv.Atoi(s)
		}
		defaultValue := 42
		expected := defaultValue
		result, err := Convert(value, convertor, defaultValue)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}
