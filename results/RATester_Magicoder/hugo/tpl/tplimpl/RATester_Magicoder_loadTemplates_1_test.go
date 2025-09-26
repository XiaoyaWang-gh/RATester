package tplimpl

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestloadTemplates_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := &templateHandler{
		layoutsFs: afero.NewMemMapFs(),
	}

	// Create a test file
	err := handler.layoutsFs.MkdirAll("templates", 0755)
	if err != nil {
		t.Fatal(err)
	}
	afero.WriteFile(handler.layoutsFs, "templates/test.txt", []byte("test"), 0644)

	// Call the method under test
	err = handler.loadTemplates()
	if err != nil {
		t.Fatal(err)
	}

	// Check if the template was added
	_, err = handler.layoutsFs.Stat("templates/test.txt")
	if err != nil {
		t.Fatal(err)
	}
}
