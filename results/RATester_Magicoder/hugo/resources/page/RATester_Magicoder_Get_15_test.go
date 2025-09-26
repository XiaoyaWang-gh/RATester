package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestGet_15(t *testing.T) {
	outputFormats := OutputFormats{
		{
			Format: output.Format{
				Name: "TestFormat",
			},
		},
	}

	t.Run("ExistingFormat", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := outputFormats.Get("TestFormat")
		if result == nil {
			t.Fatal("Expected a result, but got nil")
		}
		if result.Format.Name != "TestFormat" {
			t.Errorf("Expected format name to be 'TestFormat', but got '%s'", result.Format.Name)
		}
	})

	t.Run("NonExistingFormat", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := outputFormats.Get("NonExistingFormat")
		if result != nil {
			t.Errorf("Expected nil, but got a result")
		}
	})
}
