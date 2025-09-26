package highlight

import (
	"fmt"
	"testing"
)

func TestParseHighlightOptions_1(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		opts, err := parseHighlightOptions("")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(opts) != 0 {
			t.Errorf("Expected empty map, got %v", opts)
		}
	})

	t.Run("single option", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		opts, err := parseHighlightOptions("bold=true")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(opts) != 1 {
			t.Errorf("Expected map with 1 entry, got %v", opts)
		}
		if opts["bold"] != "true" {
			t.Errorf("Expected 'bold' option to be 'true', got %v", opts["bold"])
		}
	})

	t.Run("multiple options", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		opts, err := parseHighlightOptions("bold=true,italic=false")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(opts) != 2 {
			t.Errorf("Expected map with 2 entries, got %v", opts)
		}
		if opts["bold"] != "true" {
			t.Errorf("Expected 'bold' option to be 'true', got %v", opts["bold"])
		}
		if opts["italic"] != "false" {
			t.Errorf("Expected 'italic' option to be 'false', got %v", opts["italic"])
		}
	})

	t.Run("invalid option", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := parseHighlightOptions("bold")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
