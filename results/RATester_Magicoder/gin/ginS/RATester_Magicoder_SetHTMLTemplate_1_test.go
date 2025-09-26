package ginS

import (
	"fmt"
	"html/template"
	"testing"
)

func TestSetHTMLTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new template
	templ := template.New("test")
	templ, err := templ.Parse("<html><body>{{.}}</body></html>")
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	// Call the function under test
	SetHTMLTemplate(templ)

	// Add assertions here to verify the correctness of the function
}
