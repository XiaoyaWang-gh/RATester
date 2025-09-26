package output

import (
	"fmt"
	"testing"
)

func TestGetByNames_1(t *testing.T) {
	formats := Formats{
		{Name: "html"},
		{Name: "rss"},
		{Name: "json"},
	}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := formats.GetByNames("html", "rss")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(result) != 2 {
			t.Errorf("Expected 2 formats, got %d", len(result))
		}
		if result[0].Name != "html" || result[1].Name != "rss" {
			t.Errorf("Expected formats to be 'html' and 'rss', got %v", result)
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := formats.GetByNames("html", "non-existent")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
