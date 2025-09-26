package template

import (
	"fmt"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestOption_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new Template instance
	testTemplate := &Template{
		text: &template.Template{},
	}

	// Define the options to be passed to the Option method
	options := []string{"option1", "option2"}

	// Call the Option method
	result := testTemplate.Option(options...)

	// Assert that the result is the same as the original Template instance
	if result != testTemplate {
		t.Errorf("Expected %v, but got %v", testTemplate, result)
	}
}
