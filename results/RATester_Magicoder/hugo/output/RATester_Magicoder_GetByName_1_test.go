package output

import (
	"fmt"
	"testing"
)

func TestGetByName_1(t *testing.T) {
	formats := Formats{
		{Name: "HTML"},
		{Name: "RSS"},
	}

	t.Run("existing format", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		format, found := formats.GetByName("HTML")
		if !found {
			t.Fatal("expected format to be found")
		}
		if format.Name != "HTML" {
			t.Errorf("expected format name to be 'HTML', got %s", format.Name)
		}
	})

	t.Run("non-existing format", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, found := formats.GetByName("NonExisting")
		if found {
			t.Fatal("expected format not to be found")
		}
	})
}
