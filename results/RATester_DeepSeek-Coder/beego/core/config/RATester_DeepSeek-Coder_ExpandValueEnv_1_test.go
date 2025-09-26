package config

import (
	"fmt"
	"os"
	"testing"
)

func TestExpandValueEnv_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"no env variable", "no env", "no env"},
		{"simple env variable", "${SIMPLE}", "simple_value"},
		{"env variable with default", "${WITH_DEFAULT||default_value}", "with_value"},
		{"env variable with default at the end", "${WITH_DEFAULT_AT_END||default_value}", "with_value_at_end"},
		{"env variable with default at the beginning", "${WITH_DEFAULT_AT_BEGINNING||default_value}", "with_value_at_beginning"},
		{"env variable with default at the beginning and end", "${WITH_DEFAULT_AT_BEGINNING_AND_END||default_value}", "with_value_at_beginning_and_end"},
	}

	os.Setenv("SIMPLE", "simple_value")
	os.Setenv("WITH_VALUE", "with_value")
	os.Setenv("WITH_VALUE_AT_END", "with_value_at_end")
	os.Setenv("WITH_VALUE_AT_BEGINNING", "with_value_at_beginning")
	os.Setenv("WITH_VALUE_AT_BEGINNING_AND_END", "with_value_at_beginning_and_end")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ExpandValueEnv(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
